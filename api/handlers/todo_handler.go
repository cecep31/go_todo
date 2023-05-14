package handlers

import (
	"go_todo/api/presenter"
	"go_todo/pkg/entities"
	"go_todo/pkg/todo"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AddBook is handler/controller which creates Books in the BookShop
func AddTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Todo
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		result, err := service.InsertTodo(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(presenter.TodoSuccessResponse(result))
	}
}
func UpdateTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody entities.Todo
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(presenter.TodoErrorResponse(err))
		}

		errr := c.BodyParser(&requestBody)

		if errr != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(errr))
		}

		requestBody.ID = uint(id)

		result, err := service.UpdateTodo(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(presenter.TodoSuccessResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		err = service.RemoveTodo(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		Todo, err := service.GetTodo(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(Todo)
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
			num, errconv := strconv.Atoi(filter)
			if errconv != nil {
				return c.JSON(presenter.TodoErrorResponse(errconv))
			}
			numtouint := uint(num)
			todo, err = service.GetTodosByActivity(numtouint)
		}
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(presenter.TodosSuccessResponse(todo))
	}
}
