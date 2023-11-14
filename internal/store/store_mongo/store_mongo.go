package store_mongo

import (
	"context"
	"github.com/gorepos/usercartv2/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoStore struct {
	Store *mongo.Client
}

var db *mongo.Client

const (
	Database         = "usercart"
	ItemsCollection  = "items"
	ConnectionString = "mongodb://root:example@localhost:27017/"
)

// NewStore creates child for the Store interface
/*
There are two functions NewStore and CreateConnection, NewStore()
creates child of the interface for Store, while CreateConnection creates low level store client
*/
func NewStore() (*MongoStore, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}

	ms := &MongoStore{Store: db}
	return ms, nil
}

// CreateConnection function creates low level connection to the mongo database
func CreateConnection() (*mongo.Client, error) {
	var err error
	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to mongodb!")
	return db, nil
}

func CloseConnection() {
	err := db.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from mongodb!")
}

func (s *MongoStore) GetItems() ([]store.Item, error) {
	var items []store.Item
	collection := db.Database(Database).Collection(ItemsCollection)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var item store.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}
	return items, nil

}
