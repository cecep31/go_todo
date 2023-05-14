package main

import (
	"context"
	"fmt"
	"go_todo/api/routes"
	"go_todo/database"
	"go_todo/pkg/book"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, err := database.InitDB()
	// db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"msg": "devcode updated",
		})
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://username:password@localhost:27017/fiber").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}
