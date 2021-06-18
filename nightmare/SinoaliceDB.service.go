package nightmare

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const NightmaresUrl = "https://sinoalice.game-db.tw/package/alice_nightmares_global-en.js?"

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
