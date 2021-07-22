package nightmare

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Mongo *MongoRepository

const (
	Database = "NightmaresScheduler"
	Nightmares = "nightmare"
	URI = "MONGO_URI"
)

type Repository interface {
	SaveNightmares( nms []*Nightmare) error
	GetNightmare(nm *Nightmare) (*Nightmare, error)
	GetAll() ([]*Nightmare, error)
}
type MongoRepository struct {
	Client *mongo.Client
}

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil{
		println(err.Error())
	}
	Mongo = new(MongoRepository)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.Get(URI).(string)))
	if err !=nil {
		println(err.Error())
	}else{
		Mongo.Client = client
	}
}
func (mr *MongoRepository) SaveNightmares( nms []*Nightmare) error {
	 var (
	 	exist = new(ExistSpecification)
	 	tmp []interface{}

	 )
	 exist.Init(mr)
	 for _ , nm := range nms {
		if !exist.isSatisfiedBy(nm) {
				tmp = append(tmp, nm)
		}else{
			c :=mr.Client.Database(Database ).Collection(Nightmares)
			update := bson.M{
				"$set": nm,
			}
			_, err :=c.UpdateOne(context.Background(),bson.D{{"_id", nm.ID}}, update)
			if err != nil {
				return err
			}
		}
	}
	if len(tmp)> 0{
		_,err := mr.Client.Database(Database ).Collection(Nightmares).InsertMany(context.TODO(),tmp)
		if err != nil {
			return err
		}
	}
	return nil

}
func (mr *MongoRepository) GetNightmare(nm *Nightmare) (*Nightmare, error) {
	var tmp = new(Nightmare)
	 err := mr.Client.Database(Database).Collection(Nightmares).FindOne(context.TODO(),bson.D{{"_id", nm.ID}}).Decode(tmp)
	if err != nil {
		return tmp, err
	}

	return tmp, nil
}
func (mr *MongoRepository) GetAll() ([]*Nightmare, error) {
	var (
		tmp []*Nightmare
	)
	cursor, err := mr.Client.Database(Database).Collection(Nightmares).Find(context.TODO(),bson.D{})
	if err != nil {
		return tmp, err
	}
	if err = cursor.All(context.TODO(), &tmp); err != nil {
		return tmp, err
	}
	return tmp, nil
}
