package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/handlers"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//var user1, user2 User

var db *mongo.Client

func main() {
	log.Println("Program started...")
	//user1 = User{ID: 1, User: "user1"}
	//user2 = User{ID: 2, User: "user2"}
	//user1.Cart = []Item{}
	//user2.Cart = []Item{}

	databaseStore, err := store_mongo.NewStore()
	if err != nil {
		panic(err)
	}
	a := application.NewApplication(databaseStore)

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/healthcheck", handlers.Healthcheck)

	catalogHandler := handlers.CatalogHandler{App: a}
	v1.Get("/get_catalog", catalogHandler.GetCatalog)
	v1.Get("/items", catalogHandler.GetCatalog)
	//v1.Post("/item", handlers.AddItemHandler)
	//v1.Post("/item/:ItemID", handlers.UpdateItemHandler)
	//v1.Delete("/item/:ItemID", handlers.DeleteItemHandler)
	//v1.Get("/item/:ItemID", handlers.GetItemHandler)
	err = app.Listen(":8080")
	if err != nil {
		return
	}
}

/*func getItemFromDB(itemID primitive.ObjectID) (Item, error) {
	var item Item
	collection := db.Database("proba").Collection("Items")
	filter := bson.M{"_id": itemID}
	err := collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return Item{}, err
	}
	return item, nil
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

func createItemInDB(item Item) error {
	collection := db.Database("proba").Collection("Items")
	_, err := collection.InsertOne(context.TODO(), item)
	return err
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
*/
