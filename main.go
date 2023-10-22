package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var user1, user2 User
var item1, item2, item3, item4, item5 Item

func main() {
	item1 = NewItem(1, "Car", 1000.312)
	item2 = NewItem(2, "Pivo", 99.09)
	item3 = NewItem(3, "Kniga", 124.12)
	item4 = NewItem(4, "Iphone", 1434.1)
	item5 = NewItem(5, "Manga", 312.00)
	user1 = User{ID: 1, User: "user1"}
	user2 = User{ID: 2, User: "user2"}

	app := fiber.New()
	app.Get("/get_catalog", catalogHandler)
	app.Get("/get_user_cart", userHandler)
	app.Get("/add_item_to_cart", addHandler)
	app.Get("/remove_item_from_cart", removeHandler)
	app.Listen(":8080")
}

type User struct {
	ID      int     `json:"ID"`
	User    string  `json:"User"`
	Cart    []Item  `json:"Cart"`
	CartSum float64 `json:"CartSum"`
}
type Item struct {
	ID    int     `json:"ID"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

func removeHandler(c *fiber.Ctx) error {
	userID := c.Query("userid")
	itemID := c.Query("itemid")
	if userID == "" || itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("parametrs not valid")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user ID")
	}
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid item ID")
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("user not found")
	}
	item, err := getItemID(itemIDInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("item not found")
	}
	user.RemoveItemsFromCart(item)
	user.itemPrince()
	return c.JSON(fiber.StatusOK)
}
func addHandler(c *fiber.Ctx) error {
	userID := c.Query("userid")
	itemID := c.Query("itemid")
	if userID == "" || itemID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("parametrs not valid")
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user ID")
	}
	itemIDInt, err := strconv.Atoi(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid item ID")
	}
	var user *User
	if userIDInt == user1.ID {
		user = &user1
	} else if userIDInt == user2.ID {
		user = &user2
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("user not found")
	}
	item, err := getItemID(itemIDInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("item not found")
	}
	user.AddItemsToCart(item)
	user.itemPrince()
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
	items := []Item{item1, item2, item3, item4, item5}
	return json.NewEncoder(c).Encode(items)
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
func NewItem(ID int, Name string, Price float64) Item {
	return Item{ID, Name, Price}
}
func getItemID(itemID int) (Item, error) {
	for _, item := range []Item{item1, item2, item3, item4, item5} {
		if item.ID == itemID {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("", itemID)
}
