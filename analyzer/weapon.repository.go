package analyzer

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var WeaponRepo *WeaponRepository
const (
	Database = "Analyzer"
	Weapons = "Weapons"
	SupportSkills ="support_skill"
	URI = "MONGO_URI"
)
func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil{
		println(err.Error())
	}
	WeaponRepo = new(WeaponRepository)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.Get(URI).(string)))
	if err !=nil {
		println(err.Error())
	}else{
		WeaponRepo.Client = client
	}
}
type WeaponRepository struct {
	Client *mongo.Client
}

func (wr *WeaponRepository) SaveWeapons(weapons []*Weapon) error  {
	var tmp []interface{}
	for _, weapon := range weapons {
		tmp = append(tmp, weapon)
	}
	_ = wr.Client.Database(Database ).Collection(Weapons).Drop(context.Background())
	_,err := wr.Client.Database(Database ).Collection(Weapons).InsertMany(context.TODO(), tmp)
	if err != nil {
		return err
	}
	return nil
}
func (wr *WeaponRepository) SaveSupportSkills(supports []*SupportSkill) error  {
	var tmp []interface{}
	for _, weapon := range supports {
		tmp = append(tmp, weapon)
	}
	 _ = wr.Client.Database(Database ).Collection(SupportSkills).Drop(context.Background())
	_,err := wr.Client.Database(Database ).Collection(SupportSkills).InsertMany(context.TODO(), tmp)
	if err != nil {
		return err
	}
	return nil
}
func (wr *WeaponRepository) GetAllWeapons() ([]*Weapon, error) {
	var (
		tmp []*Weapon
	)
	cursor, err := wr.Client.Database(Database).Collection(Weapons).Find(context.Background(),bson.D{})
	if err != nil {
		return tmp, err
	}
	if err = cursor.All(context.TODO(), &tmp); err != nil {
		return tmp, err
	}
	return tmp, nil
}
func (wr *WeaponRepository) GetSupportSkillByName(name string) (SupportSkill, error) {
	tmp := new(SupportSkill)
	s :=wr.Client.Database(Database).Collection(SupportSkills)
	err := s.FindOne(context.Background(),bson.D{{"support_skill_name", name}}).Decode(tmp)
	if err != nil {
		return *tmp, err
	}

	return *tmp, nil
}
