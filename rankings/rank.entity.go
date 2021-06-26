package rankings

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rank struct {
	Name       string             `bson:"name,omitempty" json:"name"`
	Timeslot   int                `bson:"timeslot,omitempty" json:"timeslot"`
	GlobalRank int                `bson:"globalRank,omitempty" json:"globalRank"`
	TrueRank   int                `bson:"trueRank,omitempty" json:"trueRank"`
	Updated    bool               `bson:"updated,omitempty" json:"updated"`
	TsRank     int                `bson:"tsRank,omitempty" json:"tsRank"`
	Wins       int                `bson:"wins,omitempty" json:"wins"`
	Losses     int                `bson:"losses,omitempty" json:"losses"`
	Lifeforce  int                `bson:"lifeforce,omitempty" json:"lifeforce"`
}
type Ranks struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Day  int `bson:"day"`
	GC   int `bson:"gc"`
	Rank []Rank `json:"rank"`
}
