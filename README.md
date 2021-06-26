# Sinoalice_Nightmare_Scheduler_BE
Required:
* File named config.json
Ex:
```json
   {
     "MONGO_URI":"mongodb://localhost:27017",
     "PASSWORD":"your password bcrypt hash here ",
     "ORIGIN":"http://localhost:4200", 
     "PORT": "8081"
  }
```
* Go 1.12++
* MongoDB
* RSA KEY pair and the private key under the folder auth with the name "key.pem", your password must be encrypted with the public key.

To install simply use Go with the modules enabled(export GO111MODULE=on) and execute this command:
```
go get github.com/Mhoares/Sinoalice_Nightmare_Scheduler_BE
```
and in the folder of the module:
```
go build
```
and run the file that was compiled
