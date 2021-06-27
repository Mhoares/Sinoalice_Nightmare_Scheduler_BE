package nightmare

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const NightmaresUrl = "https://sinoalice.game-db.tw/package/alice_nightmares_global-en.js?"
const prefixImg = "https://sinoalice.game-db.tw/images/card/CardS"

type SinoaliceDBService struct {}

func ( sdb *SinoaliceDBService) GetNightmares( v string) ([]*Nightmare, error) {
	var (
		tmp []*Nightmare
		result struct{
			Rows []string
			Cols string
		}
	)

	params := url.Values{}
	params.Set("v", v)
	resp, errResp := http.Get(NightmaresUrl + params.Encode())
	defer resp.Body.Close()
	if errResp != nil {
		return tmp, errResp
	}
	errDecode := json.NewDecoder(resp.Body).Decode(&result)
	if errDecode  != nil {

		return tmp, errDecode
	}
	cols := strings.Split(result.Cols, "|")
	for  _,nm := range result.Rows {
		n := new(Nightmare)
		n.Init(cols ,strings.Split(nm, "|"))
		tmp = append(tmp,n)
	}
	return tmp, nil
}
func (sdb *SinoaliceDBService) GetImageDataUrl( icon string)( string, error) {
	var base64Encoding = ""
	resp, errResp := http.Get(prefixImg+ sdb.FormatedIcon(icon)+".png")
	defer resp.Body.Close()
	if errResp != nil {
		return base64Encoding, errResp
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	
	return base64Encoding, nil
	
}
func (sdb SinoaliceDBService) FormatedIcon( icon string) string  {
	n,_ := strconv.Atoi(icon)
	if n/1000 < 1 {
		return "0"+icon
	}
	return icon
}
