package analyzer

import (
	"strconv"
	"strings"
)

type Weapon struct {
	ID               int     `bson:"_id,omitempty" json:"id"`
	ResourceID       int     `bson:"resource_id" json:"resource_id"`
	Name             string  `bson:"name,omitempty" json:"name"`
	Type             string  `bson:"type" json:"type"`
	Rarity           string  `bson:"rarity" json:"rarity"`
	Cost             int     `bson:"cost" json:"cost"`
	Targets          float64 `bson:"targets" json:"targets"`
	SkillName        string  `bson:"skill_name" json:"skill_name"`
	SP               int     `bson:"sp" json:"sp"`
	SupportSkillName string  `bson:"support_skill_name" json:"support_skill_name"`
	SupportSkillTier string  `bson:"support_skill_tier" json:"support_skill_tier"`
	Damage           float64 `bson:"damage" json:"damage"`
	Recover          float64 `bson:"recover" json:"recover"`
	Patk             float64 `bson:"patk" json:"patk"`
	Matk             float64 `bson:"matk" json:"matk"`
	Pdef             float64 `bson:"pdef" json:"pdef"`
	Mdef             float64 `bson:"mdef" json:"mdef"`
	Attribute        int     `bson:"attribute" json:"attribute"`
	Evo              int     `bson:"evo" json:"evo"`
	IsInfiniteEvo    int     `bson:"is_infinite_evo" json:"is_infinite_evo"`
	MaxPatk          int     `bson:"max_patk" json:"max_patk"`
	MaxMatk          int     `bson:"max_matk" json:"max_matk"`
	MaxPdef          int     `bson:"max_pdef" json:"max_pdef"`
	MaxMdef          int     `bson:"max_mdef" json:"max_mdef"`
	AddPatk          int     `bson:"add_patk" json:"add_patk"`
	AddMatk          int     `bson:"add_matk" json:"add_matk"`
	AddMdef          int     `bson:"add_mdef" json:"add_mdef"`
	AddPdef          int     `bson:"add_pdef" json:"add_pdef"`
}

func (w *Weapon) Init(columns map[string]int, row []interface{}, detailsCol map[string]int, details [][]interface{}) {

	w.ID, _ = strconv.Atoi(row[columns["ID"]].(string))
	w.Name = row[columns["Name"]].(string)
	w.Type = row[columns["Type"]].(string)
	w.Rarity = row[columns["Rarity"]].(string)
	w.Cost, _ = strconv.Atoi(row[columns["Cost"]].(string))
	w.Targets, _ = strconv.ParseFloat(row[columns["Targets"]].(string), 32)
	w.SkillName = row[columns["Skill Name"]].(string)
	w.SP, _ = strconv.Atoi(row[columns["SP"]].(string))
	tmp := row[columns["Support Skill Name"]].(string)
	w.SupportSkillName = tmp[0 : strings.Index(tmp, "(")-1]
	w.SupportSkillTier = tmp[strings.Index(tmp, "(")+1 : strings.Index(tmp, ")")]
	if columns["Damage"] < len(row) {
		w.Damage, _ = strconv.ParseFloat(row[columns["Damage"]].(string), 32)

	}
	if columns["Recovery"] < len(row) {
		w.Recover, _ = strconv.ParseFloat(row[columns["Recovery"]].(string), 32)
	}
	if columns["P.Atk"] < len(row) {
		w.Patk, _ = strconv.ParseFloat(row[columns["P.Atk"]].(string), 32)
	}
	if columns["M.Atk"] < len(row) {
		w.Matk, _ = strconv.ParseFloat(row[columns["M.Atk"]].(string), 32)
	}
	if columns["P.Def"] < len(row) {
		w.Pdef, _ = strconv.ParseFloat(row[columns["P.Def"]].(string), 32)
	}
	if columns["M.Def"] < len(row) {
		w.Mdef, _ = strconv.ParseFloat(row[columns["M.Def"]].(string), 32)
	}
	rowDetail := w.GetRowDetail(detailsCol, details)
	if rowDetail != nil {
		w.ResourceID, _ = strconv.Atoi(rowDetail[detailsCol["resourceName"]].(string))
		w.Evo, _ = strconv.Atoi(rowDetail[detailsCol["evolutionLevel"]].(string))
		w.IsInfiniteEvo, _ = strconv.Atoi(rowDetail[detailsCol["isInfiniteEvolution"]].(string))
		w.Attribute, _ = strconv.Atoi(rowDetail[detailsCol["attribute"]].(string))
		w.MaxMdef, _ = strconv.Atoi(rowDetail[detailsCol["maxMagicDefence"]].(string))
		w.MaxPdef, _ = strconv.Atoi(rowDetail[detailsCol["maxDefence"]].(string))
		w.MaxPatk, _ = strconv.Atoi(rowDetail[detailsCol["maxAttack"]].(string))
		w.MaxMatk, _ = strconv.Atoi(rowDetail[detailsCol["maxMagicAttack"]].(string))
		w.AddPatk, _ = strconv.Atoi(rowDetail[detailsCol["addAttackValue"]].(string))
		w.AddMatk, _ = strconv.Atoi(rowDetail[detailsCol["addMagicAttackValue"]].(string))
		w.AddPdef, _ = strconv.Atoi(rowDetail[detailsCol["addDefenceValue"]].(string))
		w.AddMdef, _ = strconv.Atoi(rowDetail[detailsCol["addMagicDefenceValue"]].(string))
	}

}
func (w Weapon) GetRowDetail(detailsCol map[string]int, details [][]interface{}) []interface{} {
	for _, detail := range details {
		ID, _ := strconv.Atoi(detail[detailsCol["cardMstId"]].(string))
		if ID == w.ID {
			return detail
		}
	}
	return nil
}
