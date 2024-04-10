package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/store"
)

type AuthHandler struct {
	App *application.Application
}

var jwtSecret = []byte("your_jwt_secret_key")

func createToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var newUser store.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	existingUser, err := ah.App.S.GetUserByUsername(newUser.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to check user existence")
	}
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).SendString("User already exists")
	}

	if err := ah.App.S.CreateUser(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user")
	}

	return c.SendString("User registered successfully")
}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	user, err := ah.App.S.GetUserByUsername(credentials.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user")
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	if user.Password != credentials.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	token, err := createToken(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create token")
	}

	return c.JSON(fiber.Map{"token": token})
}
