package analyzer

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	credentialsPath = "analyzer/sinoalice-grid-analizer-ac5a82c04812.json"
	SkillSheet ="1JGtENtg-2lyyrK_5P8HjjBh-1HnVvws1s8SkwKwxjNg"
	CluesSheet="1Sy4srOxO_eyJY6YokCRgelNeMI_JHioN0BDuPnda854"
)
var BlueServ *BlueService
type ServiceAccountCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}
type BlueService struct {
	GoogleSheetApi *sheets.Service
}

func init() {
	BlueServ =new(BlueService)
	BlueServ.initGoogleSheetService()
}

func(bs *BlueService) initGoogleSheetService() {
	var credentials ServiceAccountCredentials
	ctx := context.Background()
	abs,_ :=filepath.Abs(credentialsPath)
	b, err := ioutil.ReadFile(abs)
	json.Unmarshal(b,&credentials)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config := &jwt.Config{
		Email:       credentials.ClientEmail,
		PrivateKey:  []byte(credentials.PrivateKey),
		PrivateKeyID: credentials.PrivateKeyID,
		TokenURL:     credentials.TokenURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets.readonly",
		},
	}

	client := config.Client(ctx)

	bs.GoogleSheetApi, err = sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)}
}
func (bs BlueService) GetWeapons() ( []*Weapon , error) {
	var weapons []*Weapon
	columns := map[string]int{}
	spreadsheetId := SkillSheet
	readRange := "Weapons!A:S"
	resp, err := bs.GoogleSheetApi.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		 return weapons , err
	}
	if len(resp.Values) == 0 {
		return  weapons, nil
	} else {
		columns = bs.GetColumns(resp)
		detailCol, details, errDet := bs.GetWeaponDetails()
		if errDet != nil {
			return weapons, errDet
		}
		for i, row := range resp.Values {
			if i !=0{
				w := new(Weapon)
				w.Init(columns, row,detailCol,details.Values)
				weapons = append(weapons, w)
			}
		}
	}
	return weapons, nil
}
func (bs BlueService) GetColumns(resp *sheets.ValueRange) map[string]int {
	columns := map[string]int{}
	for i, col := range resp.Values[0] {
		columns[col.(string)] = i
	}
	return columns

}
func (bs BlueService) GetWeaponDetails() (map[string]int, *sheets.ValueRange , error)   {
	columns := map[string]int{}
	spreadsheetId := SkillSheet
	readRange := "card_mst_list_v2!A:AU"
	resp, err := bs.GoogleSheetApi.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return columns, resp, err
	}
	if len(resp.Values) == 0 {
		return  columns, resp, nil
	} else {
		columns = bs.GetColumns(resp)
	}
	return columns, resp, nil

}
func (bs BlueService) GetBoostSupport(name string, r string) ([]*SupportSkill, error) {
	var support []*SupportSkill
	columns := map[string]int{}
	spreadsheetId := CluesSheet
	readRange := name+"!"+r
	resp, err := bs.GoogleSheetApi.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return support, err
	}
	if len(resp.Values) == 0 {
		return  support, nil
	} else {
		columns = bs.GetColumns(resp)
		s := new(SupportSkill)
		support = append(support, s)
		for i, row := range resp.Values {
			if i !=0{
				s.InitBoost(name,columns, row)
			}
		}
	}
	return support, nil
}
func (bs *BlueService) GetProcRates() ([]float64, error) {
	var tmp []float64
	spreadsheetId := CluesSheet
	readRange := "Proc Rate!C5:C24"
	resp, err := bs.GoogleSheetApi.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return tmp, err
	}
	if len(resp.Values) == 0 {
		return  tmp, nil
	} else {
		for _, row := range resp.Values {
			rate := row[0].(string)[0:strings.Index(row[0].(string),"%")]
			value, _ := strconv.ParseFloat(rate, 32)
			tmp =append(tmp, value)
		}
	}
	return tmp, nil
}
func (bs *BlueService) GetRawSupport() ([]*SupportSkill, error) {
	var (
		support []*SupportSkill

	)
	tmp  := map[string][][]interface{}{}
	columns := map[string]int{}
	procRates, err := bs.GetProcRates()
	spreadsheetId := SkillSheet
	readRange := "Support Multipliers!A:K"
	resp, err := bs.GoogleSheetApi.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return support, err
	}
	if len(resp.Values) == 0 {
		return  support, nil
	} else {
		columns = bs.GetColumns(resp)
		for i, row := range resp.Values {
			if i !=0{
				name := row[columns["name"]].(string)[0:strings.Index(row[columns["name"]].(string),"(")-1]
				tmp[name] = append(tmp[name], row)
			}
		}
		for name, rows := range tmp {
			s := new(SupportSkill)
			s.InitRaw(name,columns, rows, procRates)
			support = append(support, s)

		}
	}
	return support, nil

}
