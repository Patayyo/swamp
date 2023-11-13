package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	runMigration()
}

func runMigration() {
	initMongoDB()
	defer closeMongoDB()

	collectionName := "roles"

	err := db.Database("proba").Collection(collectionName).Drop(context.Background())
	if err != nil {
		log.Println("Failed to drop collection:", err)
	}

	collection := db.Database("proba").Collection(collectionName)

	document := bson.M{
		"name1": "",
		"name2": "",
		"name3": "",
		"name4": "",
		"name5": "",
	}

	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed successfully.")
}
