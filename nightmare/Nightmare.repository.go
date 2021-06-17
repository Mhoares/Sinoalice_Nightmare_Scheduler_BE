package nightmare

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Mongo *MongoRepository
type Repository interface {
	SaveNightmares( nms []*Nightmare) error
	GetNightmare(nm *Nightmare) (*Nightmare, error)
	GetAll() ([]*Nightmare, error)
}
type MongoRepository struct {
	Client *mongo.Client
}

func init() {
	Mongo = new(MongoRepository)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
		}
	}
	if len(tmp)> 0{
		_,err := mr.Client.Database("NightmaresScheduler").Collection("nightmare").InsertMany(context.TODO(),tmp)
		if err != nil {
			return err
		}
	}
	return nil

}
func (mr *MongoRepository) GetNightmare(nm *Nightmare) (*Nightmare, error) {
	var tmp = new(Nightmare)
	 err := mr.Client.Database("NightmaresScheduler").Collection("nightmare").FindOne(context.TODO(),bson.D{{"_id", nm.ID}}).Decode(tmp)
	if err != nil {
		return tmp, err
	}

	return tmp, nil
}
func (mr *MongoRepository) GetAll() ([]*Nightmare, error) {
	var (
		tmp []*Nightmare
	)
	cursor, err := mr.Client.Database("NightmaresScheduler").Collection("nightmare").Find(context.TODO(),bson.D{})
	if err != nil {
		return tmp, err
	}
	if err = cursor.All(context.TODO(), &tmp); err != nil {
		return tmp, err
	}
	return tmp, nil
}
