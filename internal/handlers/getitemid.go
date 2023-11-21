package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (ch *CatalogHandler) GetItemHandler(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	item, err := ch.App.S.GetItemByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get item"})
	}

	if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	return c.JSON(item)
}
