package main

import (
	"fmt"
	"go_todo/api/routes"
	"go_todo/database"
	"go_todo/pkg/activity"
	"go_todo/pkg/todo"
	"log"

	"github.com/joho/godotenv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	fmt.Println("init env......")
	godotenv.Load()

	fmt.Println("init db......")
	db, errdb := database.InitDB()
	// db, cancel, err := databaseConnection()
	if errdb != nil {
		log.Fatal("Database Connection Error $s", errdb)
	}
	fmt.Println("Database connection success!")
	fmt.Println("Database auto migrate")
	database.MigrateDDL(db)

	fmt.Println("create app......")
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(recover.New())

	activityRepo := activity.NewRepo(db)
	activityService := activity.NewService(activityRepo)

	todoRepo := todo.NewRepo(db)
	todeService := todo.NewService(todoRepo)

	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"msg": "devcode updated",
		})
	})
	api := app.Group("/")
	routes.ActivityRouter(api, activityService)
	routes.TodoRouter(api, todeService)
	// api := app.Group("/api")
	// routes.BookRouter(api, bookService)
	log.Fatal(app.Listen(":3030"))
}
