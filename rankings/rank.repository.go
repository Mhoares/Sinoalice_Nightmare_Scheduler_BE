package rankings

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
var RankRepo *RankRepository
const (
	Database = "RankGC"
	Collection = "Rank"
	URI = "MONGO_URI"
)
func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil{
		println(err.Error())
	}
	RankRepo = new(RankRepository)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.Get(URI).(string)))
	if err !=nil {
		println(err.Error())
	}else{
		RankRepo.Client = client
	}
}

type RankRepository struct {
	Client *mongo.Client
}

func (rr RankRepository) UpdateOrSave( rs *Ranks , day int, gc int)(bool, error)  {
	update := bson.M{
		"$set": rs,
	}

	c := rr.Client.Database(Database).Collection(Collection)
	 updated, err :=c.UpdateOne(context.Background(),bson.D{{"day", day}, {"gc", gc}}, update)
	if err != nil {
		return false, err
	}
	if  updated.MatchedCount > 0 {
		return true, nil
	}
	_, err = c.InsertOne(context.Background(), rs)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (rr RankRepository) Get( day int, gc int)(*Ranks, error)   {
	var tmp = new(Ranks)
	c := rr.Client.Database(Database).Collection(Collection)

	err := c.FindOne(context.Background(),bson.D{{"day", day}, {"gc", gc}}).Decode(tmp)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
