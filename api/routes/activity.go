package routes

import (
	"go_todo/api/handlers"
	"go_todo/pkg/activity"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func ActivityRouter(app fiber.Router, service activity.Service) {
	app.Get("/books", handlers.GetActivity(service))
	app.Post("/books", handlers.AddActivity(service))
	app.Delete("/books", handlers.RemoveActivity(service))
}
