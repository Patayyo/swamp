package main

import (
	"context"
	"github.com/gorepos/usercartv2/internal/store"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func main() {
	runMigration()
}

func runMigration() {
	log.Println("Run migration...")
	db, err := store_mongo.CreateConnection()
	if err != nil {
		panic(err)
	}
	defer store_mongo.CloseConnection()

	err = db.Database(store_mongo.Database).Drop(context.Background())
	if err != nil {
		panic(err)
	}

	collection := db.Database(store_mongo.Database).Collection(store_mongo.ItemsCollection)

	documents := []interface{}{
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Car",
			Price: 100,
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Airplane",
			Price: 100000,
		},
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		panic(err)
	}

	log.Println("Migration completed successfully.")
}
