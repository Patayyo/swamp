package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
)

type CurrencyHandler struct {
	App *application.Application
}

type CurrencyRequest struct {
	Username string  `json:"username"`
	Amount   float64 `json:"amount"`
}

func (ch *CurrencyHandler) AddCurrency(c *fiber.Ctx) error {
	var req CurrencyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ch.App.S.UpdateBalance(req.Username, req.Amount); err != nil {
		log.Printf("Error adding currency: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add currency")
	}

	return c.SendString("Currency added successfully")
}

func (ch *CurrencyHandler) DeductCurrency(c *fiber.Ctx) error {
	var req CurrencyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	if err := ch.App.S.UpdateBalance(req.Username, -req.Amount); err != nil {
		log.Printf("Error deducting currency: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to deduct currency")
	}

	return c.SendString("Currency deducted successfully")
}

func (ch *CurrencyHandler) GetBalance(c *fiber.Ctx) error {
	token := extractTokenFromRequest(c)
	username, err := extractUserIDFromToken(token)
	if err != nil {
		log.Printf("Error extracting user ID from token: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	user, err := ch.App.S.GetUserByUsername(username)
	if err != nil || user == nil {
		log.Printf("Error retrieving user: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user")
	}

	return c.JSON(fiber.Map{"balance": user.Balance})
}
