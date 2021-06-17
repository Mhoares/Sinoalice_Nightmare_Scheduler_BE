package nightmare

import (
	"strconv"
)

const (
	ID = "ID"
	Icon = "Icon"
	Rarity = "Rarity"
	Attribute = "Attribute"
	NameEN ="NameEN"
	GvgSkillEN = "GvgSkillEN"
	GvgSkillDur = "GvgSkillDur"
	GvgSkillLead = "GvgSkillLead"

)

type Nightmare struct {
	ID int `bson:"_id,omitempty"`
	Icon string `bson:"icon,omitempty"`
	Attribute int `bson:"attribute,omitempty"`
	Rarity int	`bson:"rarity,omitempty"`
	NameEN string `bson:"name_en,omitempty"`
	GvgSkillEN string	`bson:"gvg_skill_en,omitempty"`
	GvgSkillDur int	`bson:"gvg_skill_dur,omitempty"`
	GvgSkillLead int	`bson:"gvg_skill_lead,omitempty"`
}
func(n *Nightmare) Init(columns []string, row []string){
	n.ID, _ = strconv.Atoi(row[n.indexOf(ID, columns)])
	n.Icon = row[n.indexOf(Icon, columns)]
	n.Attribute, _ = strconv.Atoi( row[n.indexOf(Attribute,columns)])
	n.Rarity, _ = strconv.Atoi( row[n.indexOf(Rarity,columns)])
	n.NameEN = row[n.indexOf(NameEN,columns)]
	n.GvgSkillEN = row[n.indexOf(GvgSkillEN,columns)]
	n.GvgSkillDur, _ =   strconv.Atoi( row[n.indexOf(GvgSkillDur,columns)])
	n.GvgSkillLead, _ =   strconv.Atoi( row[n.indexOf(GvgSkillLead,columns)])
}
func(n *Nightmare) indexOf( column string, columns []string) int{
	for i, c := range columns{
		if c == column {
			return i
		}
	}
	return -1
}
