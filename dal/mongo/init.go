package mongo

import (
	"context"
	"github.com/yuanyangen/trader1024/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var Trader1024Db *mongo.Database

func init() {
	var err error
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(config.GetMasterMongoUri()))
	if err != nil {
		panic(err)
	}
	Trader1024Db = MongoClient.Database("trader_1024")
}
