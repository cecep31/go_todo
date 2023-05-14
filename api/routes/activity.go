package routes

import (
	"go_todo/api/handlers"
	"go_todo/pkg/activity"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func ActivityRouter(app fiber.Router, service activity.Service) {
	app.Get("/activity-groups", handlers.GetActivities(service))
	app.Get("/activity-groups/:id", handlers.GetActivity(service))
	app.Post("/activity-groups", handlers.AddActivity(service))
	app.Delete("/activity-groups/:id", handlers.RemoveActivity(service))
}
