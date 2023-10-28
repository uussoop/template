package history

import (
	"context"
	"test/database"
	"test/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Messages struct{ M []utils.Message }

var historyCollection *mongo.Collection

func getHistoryCollection() *mongo.Collection {
	if historyCollection == nil {
		historyCollection = database.DB.Database("test").Collection("test")
	}
	return historyCollection
}

func Insert(m *[]utils.Message) (err error) {
	Mess := Messages{M: *m}
	_, err = getHistoryCollection().InsertOne(context.Background(), Mess)
	return
}
func Read() ([]utils.Message, error) {

	var osme Messages
	err := getHistoryCollection().FindOne(context.Background(), bson.M{}).Decode(&osme)
	if err != nil {
		return nil, err
	}

	return osme.M, nil
}
