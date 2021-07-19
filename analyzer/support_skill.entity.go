package analyzer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

type SupportSkill struct {
	ID               primitive.ObjectID                   `bson:"_id,omitempty" json:"-"`
	SupportSkillName string                               `bson:"support_skill_name" json:"support_skill_name"`
	SupportSkillTier map[string][]SupportSkillDescription `bson:"support_skill_tier" json:"support_skill_tier"`
}

type SupportSkillDescription struct {
	Rate  float64            `bson:"rate" json:"rate"`
	Value map[string]float64 `bson:"value" json:"value"`
}

func (ss *SupportSkill) InitBoost(name string, cols map[string]int, row []interface{}) {
	var (
		value float64
	)

	if len(ss.SupportSkillName) == 0 {
		ss.SupportSkillName = name
	}
	for i, col := range cols {
		tmp := row[col].(string)[0:strings.Index(row[col].(string), "%")]
		 if i != "Rate" {
			if ss.SupportSkillTier == nil {
				ss.SupportSkillTier = map[string][]SupportSkillDescription{}
			}
			value, _ = strconv.ParseFloat(tmp, 32)
			ss.SupportSkillTier[i] = append(ss.SupportSkillTier[i], SupportSkillDescription{Rate: ss.getRate(cols, row), Value: map[string]float64{
				"boost": value,
			}})
		}

	}
}
func (ss *SupportSkill) getRate(cols map[string]int ,row []interface{}) float64 {
	tmp := row[cols["Rate"]].(string)[0:strings.Index(row[cols["Rate"]].(string), "%")]
	rate, _ := strconv.ParseFloat(tmp, 32)
	return rate
}
func (ss *SupportSkill) InitRaw(name string, cols map[string]int, rows [][]interface{}, procRates []float64) {

	ss.SupportSkillName = name
	ss.SupportSkillTier = map[string][]SupportSkillDescription{}
	value := map[string]float64{}
	for _, row := range rows {
		rawName := row[cols["name"]].(string)
		tier := rawName[strings.Index(rawName, "(")+1 : strings.Index(rawName, ")")]
		if cols["P.Atk"] < len(row) {
			value["P.Atk"], _ = strconv.ParseFloat(row[cols["P.Atk"]].(string),32)
		}
		if cols["M.Atk"] < len(row) {
			value["M.Atk"], _ = strconv.ParseFloat(row[cols["M.Atk"]].(string),32)
		}
		if cols["P.Def"] < len(row) {
			value["P.Def"], _ = strconv.ParseFloat(row[cols["P.Def"]].(string),32)
		}
		if cols["M.Def"] < len(row) {
			value["M.Def"], _ = strconv.ParseFloat(row[cols["M.Def"]].(string),32)
		}
		for _ ,rate := range procRates {
			ss.SupportSkillTier[tier] = append(ss.SupportSkillTier[tier],SupportSkillDescription{Rate: rate, Value: value})
		}
	}

}
