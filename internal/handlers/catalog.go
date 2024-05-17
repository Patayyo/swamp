package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/store"
)

type CatalogHandler struct {
	App *application.Application
}

func (ch *CatalogHandler) GetCatalog(c *fiber.Ctx) error {
	items, err := ch.App.S.GetItems()
	if err != nil {
		log.Printf("Error getting catalog items: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("500")
	}
	return c.JSON(items)
}

func (ch *CatalogHandler) AddItemToCart(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	token := extractTokenFromRequest(c)
	userID, err := extractUserIDFromToken(token)
	if err != nil {
		log.Printf("Error extracting user ID from token: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	if err := ch.App.S.AddItemToCart(userID, itemID); err != nil {
		log.Printf("Error adding item to cart: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add item to cart")
	}

	log.Printf("Item added to cart successfully")
	return c.SendString("Item added to cart successfully")
}

func (ch *CatalogHandler) RemoveItemFromCart(c *fiber.Ctx) error {
	itemID := c.Params("ItemID")

	token := extractTokenFromRequest(c)
	userID, err := extractUserIDFromToken(token)
	if err != nil {
		log.Printf("Error extracting user ID from token: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	if err := ch.App.S.RemoveItemFromCart(userID, itemID); err != nil {
		log.Printf("Error removing item from cart: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to remove item from cart")
	}

	log.Printf("Item removed from cart successfully")
	return c.SendString("Item removed from cart successfully")
}

func (ch *CatalogHandler) GetCart(c *fiber.Ctx) error {
	token := extractTokenFromRequest(c)
	userID, err := extractUserIDFromToken(token)
	if err != nil {
		log.Printf("Error extracting user ID from token: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	cart, err := ch.App.S.GetCart(userID)
	if err != nil {
		log.Printf("Error getting cart: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get cart")
	}

	if cart == nil {
		log.Printf("Cart not found, returning empty cart for user: %s", userID)
		cart = []store.Item{}
	}

	return c.JSON(cart)
}
