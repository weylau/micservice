package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"user-edge-service/app/config"
	"user-edge-service/app/loger"
	"time"
)

type Mongo struct {
	conn *mongo.Client
}

func Default() *Mongo {
	mg := &Mongo{}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Configs.MongoHost))
	if err != nil {
		loger.Default().Error("mongo conn error:", err.Error())
		panic(err.Error())
	}
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		loger.Default().Error("mongo ping error:", err.Error())
		panic(err.Error())
	}
	mg.conn = client
	return mg
}

func (this *Mongo) GetConn() *mongo.Client {
	return this.conn
}
