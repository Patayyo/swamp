package main

import (
	"context"

	//"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var user1, user2 User
var item1, item2, item3, item4, item5 Item
var db *mongo.Client

func main() {
	/*item1 = NewItem(1, "Car", 1000.312)
	item2 = NewItem(2, "Pivo", 99.09)
	item3 = NewItem(3, "Kniga", 124.12)
	item4 = NewItem(4, "Iphone", 1434.1)
	item5 = NewItem(5, "Manga", 312.00)*/
	var err error
	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	items, err := loadItemsFromDB()
	if err != nil {
		log.Fatal(err)
	}

	user1 = User{ID: 1, User: "user1"}
	user2 = User{ID: 2, User: "user2"}
	user1.Cart = []Item{}
	user2.Cart = []Item{}
	var itemsList []Item
	for _, item := range items {
		itemsList = append(itemsList, item)
	}

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/healthcheck", healthcheck)
	v1.Get("/get_catalog", catalogHandler)
	v1.Get("/items", catalogHandler)
	//v1.Get("/get_user_cart", userHandler)
	//v1.Get("/add_item_to_cart", addHandler)
	v1.Post("/item", addItemHandler)
	v1.Post("/item/:ItemID", updateItemHandler)
	v1.Delete("/item/:ItemID", deleteItemHandler)
	v1.Get("/item/:ItemID", getItemHandler)
	//v1.Get("/remove_item_from_cart", removeHandler)

	fmt.Println("Connected to MongoDB!")
	defer db.Disconnect(context.TODO())
	app.Listen(":8080")
}

type User struct {
	ID      int     `json:"ID"`
	User    string  `json:"User"`
	Cart    []Item  `json:"Cart"`
	CartSum float64 `json:"CartSum"`
}
type Item struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"Name"`
	Price float64            `json:"Price"`
}
type NewItemInput struct {
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

func getItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	item, err := getItemFromDB(itemObjectID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("404")
	}
	return c.JSON(item)
}
func getItemFromDB(itemID primitive.ObjectID) (Item, error) {
	var item Item
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"_id": itemID}
	err := collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}
func deleteItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"_id": itemObjectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func updateItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")
	if itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	itemObjectID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	var updatedItem Item
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	err = updateItemInDB(itemObjectID, updatedItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(updatedItem)
}
func updateItemInDB(itemID primitive.ObjectID, updatedItem Item) error {
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"_id": itemID}
	update := bson.M{"$set": bson.M{"Name": updatedItem.Name, "Price": updatedItem.Price}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func addItemHandler(c *fiber.Ctx) error {
	var input NewItemInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("400")
	}
	newItem := Item{
		Name:  input.Name,
		Price: input.Price,
	}
	if err := createItemInDB(newItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(newItem)
}
func createItemInDB(item Item) error {
	collection := db.Database("proba").Collection("Items")
	_, err := collection.InsertOne(context.TODO(), item)
	return err
}
func healthcheck(c *fiber.Ctx) error {
	return c.JSON(fiber.StatusOK)
}

func userHandler(c *fiber.Ctx) error {
	userID := c.Query("userid")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("user not found")
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}
	return c.JSON(user)
}
func catalogHandler(c *fiber.Ctx) error {
	items, err := loadItemsFromDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(items)
}
func loadItemsFromDB() ([]Item, error) {
	var items []Item
	collection := db.Database("proba").Collection("Items")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var item Item
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
func updateItemsInDB(items []Item) error {
	collection := db.Database("proba").Collection("Items")
	_, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	documents := make([]interface{}, len(items))
	for i, item := range items {
		documents[i] = item
	}
	_, err = collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}
	return nil
}
func generateNewItemID() int {
	for {
		newID := rand.Int63()
		if newID < 0 {
			newID = newID + math.MaxInt32
		}
		if !isItemIDUsed(int(newID)) {
			return int(newID)
		}
	}
}
func isItemIDUsed(newID int) bool {
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"ID": newID}
	var existingItem Item
	err := collection.FindOne(context.TODO(), filter).Decode(&existingItem)
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Println("shit", err)
		return true
	}
	return true
}
func (u *User) itemPrince() {
	sum := 0.0
	for _, item := range u.Cart {
		sum += item.Price
	}
	u.CartSum = sum
}
func (u *User) AddItemsToCart(item Item) {
	u.Cart = append(u.Cart, item)
}
func (u *User) ShowUserCart() {

}
func (u *User) RemoveItemsFromCart(item Item) {
	for i, v := range u.Cart {
		if v == item {
			copy(u.Cart[i:], u.Cart[i+1:])
			u.Cart = u.Cart[:len(u.Cart)-1]
		}
	}
	return
}
