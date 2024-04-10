package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gorepos/usercartv2/internal/application"
	"github.com/gorepos/usercartv2/internal/handlers"
	"github.com/gorepos/usercartv2/internal/store/store_mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func main() {
	log.Println("Program started...")

	databaseStore, err := store_mongo.NewStore()
	if err != nil {
		panic(err)
	}
	a := application.NewApplication(databaseStore)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
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
	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
