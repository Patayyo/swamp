package main

import (
	"context"
	"log"

	"github.com/gorepos/usercartv2/internal/store"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			Name:  "Azuki",
			Price: 6.18,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Chocola",
			Price: 5.86,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cinnamon",
			Price: 7.62,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Coconut",
			Price: 5.15,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Maple",
			Price: 4.51,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Minaduki Family",
			Price: 5.15,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Shigure",
			Price: 4.43,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Vanilla",
			Price: 6.83,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cuphead",
			Price: 2.43,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Mugman",
			Price: 2.10,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Elder Kettle",
			Price: 2.24,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Legendary Chalice",
			Price: 2.31,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "The Devil",
			Price: 2.57,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "King Dice",
			Price: 2.60,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Hilda Berg",
			Price: 2.91,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Cala Maria",
			Price: 2.24,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "AAA",
			Price: 2.77,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Large Arachnoid",
			Price: 1.77,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Biomechanoid",
			Price: 2.27,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Gnaar",
			Price: 2.16,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Khnum",
			Price: 2.16,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Octanian",
			Price: 2.26,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Processed",
			Price: 2.10,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Draconian Pyromaniac",
			Price: 2.23,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Biker Sam",
			Price: 2.50,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Scrapjack",
			Price: 2.22,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Ugh Zan",
			Price: 2.30,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Sirian Werebull",
			Price: 2.27,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Beheaded Kamikaze",
			Price: 2.24,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Kleer Skeleton",
			Price: 2.13,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Lord Achriman",
			Price: 2.30,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Balkan",
			Price: 2,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "Anarchist",
			Price: 2,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "IDF",
			Price: 2,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "FBI",
			Price: 1.72,
			Type:  "20",
		},
		store.Item{
			ID:    primitive.ObjectID{},
			Name:  "SWAT",
			Price: 1.73,
			Type:  "20",
		},
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		panic(err)
	}

	log.Println("Migration completed successfully.")
}
