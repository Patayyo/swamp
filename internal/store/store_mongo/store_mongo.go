package store_mongo

import (
	"context"
	"errors"
	"log"

	"github.com/gorepos/usercartv2/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Store *mongo.Client
}

var db *mongo.Client

const (
	Database         = "usercart"
	ItemsCollection  = "items"
	UsersCollection  = "user"
	CartCollection   = "cart"
	ConnectionString = "mongodb://root:example@mongo:27017/"
)

func NewStore() (*MongoStore, error) {
	db, err := CreateConnection()
	if err != nil {
		return nil, err
	}

	ms := &MongoStore{Store: db}
	return ms, nil
}

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
	collection := s.Store.Database("usercart").Collection("items")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("Error occurred while finding items: %v\n", err)
		return nil, err
	}

	var items []store.Item
	err = cursor.All(context.TODO(), &items)
	if err != nil {
		log.Printf("Error occurred while decoding items: %v\n", err)
		return nil, err
	}

	return items, nil
}

func (s *MongoStore) AddItem(item store.Item) error {
	collection := s.Store.Database("usercart").Collection("items")
	_, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Printf("Error occurred while inserting item: %v\n", err)
		return err
	}

	return nil
}

func (s *MongoStore) GetItemByID(id string) (*store.Item, error) {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var item store.Item
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (s *MongoStore) UpdateItem(id string, updatedItem store.Item) error {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":  updatedItem.Name,
		"price": updatedItem.Price,
		"type":  updatedItem.Type,
	}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) DeleteItem(id string) error {
	collection := s.Store.Database(Database).Collection(ItemsCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoStore) CreateUser(user store.User) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	user.Cart = []store.Item{}

	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (s *MongoStore) GetUserByUsername(username string) (*store.User, error) {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	var user store.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Error getting user by username %s: %v", username, err)
		return nil, err
	}

	return &user, nil
}

func (s *MongoStore) AddItemToCart(userID, itemID string) error {
	collection := s.Store.Database(Database).Collection(CartCollection)

	item, err := s.GetItemByID(itemID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$push": bson.M{"items": item}},
		options.Update().SetUpsert(true),
	)
	return err
}

func (s *MongoStore) RemoveItemFromCart(userID, itemID string) error {
	collection := s.Store.Database(Database).Collection(CartCollection)
	objectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"items": bson.M{"_id": objectID}}},
	)
	return err
}

func (s *MongoStore) GetCart(userID string) ([]store.Item, error) {
	collection := s.Store.Database(Database).Collection(CartCollection)
	filter := bson.M{"_id": userID}

	var cart struct {
		Items []store.Item `bson:"items"`
	}

	err := collection.FindOne(context.Background(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return cart.Items, nil
}

func (s *MongoStore) CreateCart(userID string) error {
	collection := s.Store.Database(Database).Collection(CartCollection)
	_, err := collection.InsertOne(context.Background(), bson.M{
		"_id":   userID,
		"items": []store.Item{},
	})
	return err
}

func (s *MongoStore) GetUserByID(userID string) (*store.User, error) {
	if userID == "" {
		return nil, errors.New("userID is empty")
	}

	if len(userID) != 24 {
		return nil, errors.New("userID is not a valid ObjectID")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	collection := s.Store.Database(Database).Collection(UsersCollection)
	var user store.User
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Error getting user by ID %s: %v", userID, err)
		return nil, err
	}

	return &user, nil
}

func (s *MongoStore) DeleteUser(id string) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoStore) GetUsers() ([]store.User, error) {
	collection := s.Store.Database(Database).Collection(UsersCollection)

	var users []store.User
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user store.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoStore) UpdateBalance(username string, amount float64) error {
	collection := s.Store.Database(Database).Collection(UsersCollection)
	filter := bson.M{"username": username}
	update := bson.M{"$inc": bson.M{"balance": amount}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
