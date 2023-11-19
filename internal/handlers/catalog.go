package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
)

type CatalogHandler struct {
	App *application.Application
}

func (ch *CatalogHandler) GetCatalog(c *fiber.Ctx) error {
	items, err := ch.App.S.GetItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(items)
}
