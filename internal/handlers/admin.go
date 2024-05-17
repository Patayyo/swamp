package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type AdminHandler struct {
	App *application.Application
}

func (ah *AdminHandler) GetItems(c *fiber.Ctx) error {
	items, err := ah.App.S.GetItems()
	if err != nil {
		log.Printf("Error getting items: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get items")
	}
	return c.JSON(items)
}

func (ah *AdminHandler) CreateItem(c *fiber.Ctx) error {
	var newItem store.Item
	if err := c.BodyParser(&newItem); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ah.App.S.AddItem(newItem); err != nil {
		log.Printf("Error adding item: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add item")
	}

	log.Printf("Item added successfully: %+v", newItem)
	return c.JSON(newItem)
}

func (ah *AdminHandler) GetItem(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	item, err := ah.App.S.GetItemByID(itemID)
	if err != nil {
		log.Printf("Error getting item: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get item")
	}
	if item == nil {
		return c.Status(fiber.StatusNotFound).SendString("Item not found")
	}

	return c.JSON(item)
}

func (ah *AdminHandler) UpdateItem(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	var updatedItem store.Item
	if err := c.BodyParser(&updatedItem); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ah.App.S.UpdateItem(itemID, updatedItem); err != nil {
		log.Printf("Error updating item: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update item")
	}

	log.Printf("Item updated successfully: %+v", updatedItem)
	return c.JSON(updatedItem)
}

func (ah *AdminHandler) DeleteItem(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	if err := ah.App.S.DeleteItem(itemID); err != nil {
		log.Printf("Error deleting item: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete item")
	}

	log.Printf("Item deleted successfully")
	return c.SendString("Item deleted successfully")
}

func (ah *AdminHandler) GetUsers(c *fiber.Ctx) error {
	users, err := ah.App.S.GetUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get users")
	}
	return c.JSON(users)
}

func (ah *AdminHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("UserID")
	log.Printf("Received userID: %s", userID)

	if !bson.IsObjectIdHex(userID) {
		log.Printf("Invalid userID format: %s", userID)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID format")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error converting userID to ObjectID: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid userID")
	}
	log.Printf("Converted ObjectID: %s", objectID.Hex())

	user, err := ah.App.S.GetUserByID(objectID.Hex())
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user")
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

func (ah *AdminHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("UserID")

	if err := ah.App.S.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete user")
	}

	log.Printf("User deleted successfully")
	return c.SendString("User deleted successfully")
}
