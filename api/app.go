package main

import (
	"fmt"
	"go_todo/api/routes"
	"go_todo/database"
	"go_todo/pkg/activity"
	"log"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("init env......")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("init db......")
	db, err := database.InitDB()
	// db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")

	fmt.Println("Database auto migrate")
	database.MigrateDDL(db)

	fmt.Println("create app......")
	app := fiber.New()
	activityRepo := activity.NewRepo(db)
	activityService := activity.NewService(activityRepo)

	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"msg": "devcode updated",
		})
	})
	api := app.Group("/")
	routes.ActivityRouter(api, activityService)
	// api := app.Group("/api")
	// routes.BookRouter(api, bookService)
	log.Fatal(app.Listen(":8090"))
}
