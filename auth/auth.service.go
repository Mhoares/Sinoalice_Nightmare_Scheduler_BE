package auth

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "encoding/base64"
  "encoding/pem"
  "github.com/spf13/viper"
  "golang.org/x/crypto/bcrypt"

  "io/ioutil"
  "path/filepath"
)


const(
  password = "PASSWORD"
  privateKeyPath = "auth/key.pem"
)

type Service struct {
   PrivateKey *rsa.PrivateKey
}

func (as *Service)Init()  error {
  privateKeyAbs, _ := filepath.Abs(privateKeyPath)
  privateBytes, err := ioutil.ReadFile(privateKeyAbs)
  privatePem , _ := pem.Decode(privateBytes)

  as.PrivateKey, err =x509.ParsePKCS1PrivateKey(privatePem.Bytes)
  if err != nil {
    return err
  }
  return nil
}
func (as *Service) CheckPassword( pwd string)(bool, error)  {
  viper.SetConfigFile("config.json")
  if err := viper.ReadInConfig(); err != nil{
    println(err.Error())
  }
  data, err64 := base64.StdEncoding.DecodeString(pwd)
  if err64 != nil {
    return false, err64
  }
  dec, err := rsa.DecryptPKCS1v15(rand.Reader,as.PrivateKey,data)
  if err != nil {
    return false, err
  }
  err = bcrypt.CompareHashAndPassword([]byte(viper.Get(password).(string)), dec)
  if err != nil {
    return false, err
  }
  return true, nil
}
