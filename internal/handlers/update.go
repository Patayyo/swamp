package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/store"
)

func (ch *CatalogHandler) UpdateItemHandler(c *fiber.Ctx) error {

	itemID := c.Params("ItemID")

	var updatedItem store.Item
	if err := c.BodyParser(&updatedItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid requset format"})
	}

	if err := ch.App.S.UpdateItem(itemID, updatedItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update item"})
	}

	return c.JSON(fiber.Map{"message": "Item updated successfully"})
}
