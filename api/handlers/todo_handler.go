package handlers

import (
	"fmt"
	"go_todo/api/presenter"
	"go_todo/pkg/entities"
	"go_todo/pkg/todo"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddBook is handler/controller which creates Books in the BookShop
func AddTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(entities.Todo)
		errbody := c.BodyParser(&requestBody)
		if errbody != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TodoErrorResponse(errbody))
		}
		requestBody.ID = uuid.New()
		result, err := service.InsertTodo(requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.TodoSuccessResponse(result))
	}
}
func UpdateTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Todo
		id := c.Params("id")
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}

		errbody := c.BodyParser(&requestBody)

		if errbody != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.TodoErrorResponse(errbody))
		}

		requestBody.ID = idparam

		result, err := service.UpdateTodo(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}
		return c.JSON(presenter.TodoSuccessResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}
		err := service.RemoveTodo(idparam)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "Success",
			"message": "sucess remove",
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}
		todo, errservice := service.GetTodo(idparam)
		if errservice != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %v Not Found", id),
			})
		}
		return c.JSON(presenter.TodoSuccessResponse(todo))
	}
}

func GetTodos(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := c.FormValue("activity_group_id")

		var todo *[]entities.Todo
		var err error

		if filter == "" {
			todo, err = service.GetTodos()
		} else {
			numtouint, erruuid := uuid.Parse(filter)
			if erruuid != nil {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"status":  "Success",
					"message": "Success",
					"data":    todo,
				})
			}
			todo, err = service.GetTodosByActivity(numtouint)
		}
		if err != nil {
			return c.JSON(presenter.TodosSuccessResponse(todo))
		}
		return c.JSON(presenter.TodosSuccessResponse(todo))
	}
}
