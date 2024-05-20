package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/handlers"
	"github.com/gorepos/usercartv2/internal/store"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func main() {
	log.Println("Program started...")

	databaseStore, err := store_mongo.NewStore()
	if err != nil {
		log.Fatalf("Error creating database store: %v", err)
	}

	if err := databaseStore.UpdateAllUsersRole(); err != nil {
		log.Fatalf("Error updating users' roles: %v", err)
	}
	log.Println("User roles updated successfully")

	adminUser := store.User{
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
		Balance:  1000,
	}

	if err := databaseStore.CreateUser(adminUser); err != nil {
		log.Fatalf("Error creating admin user: %v", err)
	}

	log.Println("Admin user created successfully")
	a := application.NewApplication(databaseStore)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE ,PUT",
	}))

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/healthcheck", handlers.Healthcheck)

	authHandler := handlers.AuthHandler{App: a}
	v1.Post("/reg", authHandler.Register)
	v1.Post("/login", authHandler.Login)

	catalogHandler := handlers.CatalogHandler{App: a}
	v1.Get("/get_catalog", catalogHandler.GetCatalog)
	v1.Get("/items", catalogHandler.GetCatalog)
	v1.Post("/item", catalogHandler.AddItemHandler)
	v1.Post("/item/:ItemID", catalogHandler.UpdateItemHandler)
	v1.Delete("/item/:ItemID", catalogHandler.DeleteItemHandler)
	v1.Get("/item/:ItemID", catalogHandler.GetItemHandler)
	v1.Post("/cart/add/:ItemID", catalogHandler.AddItemToCart)
	v1.Post("/cart/remove/:ItemID", catalogHandler.RemoveItemFromCart)
	v1.Get("/cart", catalogHandler.GetCart)

	currencyHandler := handlers.CurrencyHandler{App: a}
	v1.Get("/currency/balance", currencyHandler.GetBalance)

	adminHandler := handlers.AdminHandler{App: a}
	adminMiddleware := authHandler.AdminMiddleware
	admin := v1.Group("/admin", adminMiddleware)

	admin.Get("/items", adminHandler.GetItems)
	admin.Post("/items", adminHandler.CreateItem)
	admin.Get("/items/:ItemID", adminHandler.GetItem)
	admin.Put("/items/:ItemID", adminHandler.UpdateItem)
	admin.Delete("/items/:ItemID", adminHandler.DeleteItem)

	admin.Get("/users", adminHandler.GetUsers)
	admin.Get("/users/:UserID", adminHandler.GetUser)
	admin.Delete("/users/:UserID", adminHandler.DeleteUser)

	admin.Post("/currency/add", currencyHandler.AddCurrency)
	admin.Post("/currency/deduct", currencyHandler.DeductCurrency)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
