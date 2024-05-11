package handlers

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

func createToken(username, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role

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
		Username string `json:"username" bson:"username,omitempty"`
		Password string `json:"password" bson:"password,omitempty"`
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

	token, err := createToken(user.Username, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create token")
	}

	return c.JSON(fiber.Map{"token": token})
}

func extractTokenFromRequest(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		token := c.Cookies("jwt")
		return token
	}

	return strings.Replace(authHeader, "Bearer ", "", 1)
}

func extractUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username not found in token")
	}

	return username, nil
}

func (ah *AuthHandler) AdminMiddleware(c *fiber.Ctx) error {
	token := extractTokenFromRequest(c)
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	username, err := extractUserIDFromToken(token)
	if err != nil {
		log.Printf("Error extracting user ID from token: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	user, err := ah.App.S.GetUserByUsername(username)
	if err != nil || user == nil {
		log.Printf("Error retrieving user: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	}

	return c.Next()
}
