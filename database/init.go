package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *mongo.Client
)

func InitDB() {
	op := options.Client()
	op.ApplyURI("mongodb://localhost:27017")
	op.Auth = &options.Credential{
		Username: "root",
		Password: "root",
	}
	var err error
	DB, err = mongo.Connect(context.TODO(), op)

	if err != nil {
		panic(err)
	}
	err = DB.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
