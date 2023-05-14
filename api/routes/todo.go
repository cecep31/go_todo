package routes

import (
	"go_todo/api/handlers"
	"go_todo/pkg/todo"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is the Router for GoFiber App
func TodoRouter(app fiber.Router, service todo.Service) {
	app.Get("/todo-items", handlers.GetTodos(service))
	app.Get("/todo-items/:id", handlers.GetTodo(service))
	app.Post("/todo-items", handlers.AddTodo(service))
	app.Patch("/todo-items/:id", handlers.UpdateTodo(service))
	app.Delete("/todo-items/:id", handlers.RemoveTodo(service))
}
