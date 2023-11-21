package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/store"
)

func (ch *CatalogHandler) AddItemHandler(c *fiber.Ctx) error {

	var newItem store.Item
	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ch.App.S.AddItem(newItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add item to the catalog")
	}

	return c.SendString("Item added to the catalog successfully")
}
