package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/store"
)

func (ch *CatalogHandler) DeleteItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	err := ch.App.S.DeleteItem(itemID)
	if err != nil {
		log.Printf("Error deleting item with ID %s: %v", itemID, err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Serrver Error")
	}

	log.Printf("Item with ID %s deleted successfully", itemID)
	return c.SendString("Item deleted successfully")
}

func (ch *CatalogHandler) AddItemHandler(c *fiber.Ctx) error {
	var newItem store.Item

	if err := c.BodyParser(&newItem); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ch.App.S.AddItem(newItem); err != nil {
		log.Printf("Error adding item: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add item to the catalog")
	}

	log.Printf("Item added to the catalog successfully: %+v", newItem)
	return c.SendString("Item added to the catalog successfully")
}

func (ch *CatalogHandler) GetItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	item, err := ch.App.S.GetItemByID(itemID)
	if err != nil {
		log.Printf("Error getting item with ID %s: %v", itemID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get item"})
	}

	if item == nil {
		log.Printf("Item with ID %s not found", itemID)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	log.Printf("Item with ID %s retrieved successfully: %+v", itemID, item)
	return c.JSON(item)
}

func (ch *CatalogHandler) UpdateItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	var updatedItem store.Item
	if err := c.BodyParser(&updatedItem); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid requset format"})
	}

	if err := ch.App.S.UpdateItem(itemID, updatedItem); err != nil {
		log.Printf("Error updating item with ID %s: %v", itemID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update item"})
	}

	log.Printf("Item with ID %s updated successfully: %+v", itemID, updatedItem)
	return c.JSON(fiber.Map{"message": "Item updated successfully"})
}
