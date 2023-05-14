package presenter

import (
	"go_todo/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler
type Todo struct {
	ID        uint `gorm:"primarykey;column:todo_id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string `json:"title"`
	Is_active bool   `json:"is_active"`
	Priority  string `json:"priority"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handle
func TodoSuccessResponse(data *entities.Todo) *fiber.Map {
	Todo := Todo{
		ID:        data.ID,
		Title:     data.Title,
		Is_active: data.Is_active,
		Priority:  data.Priority,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    Todo,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func TodosSuccessResponse(Todos *[]entities.Todo) *fiber.Map {
	return &fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    Todos,
	}
}

func TodoSuccessCreateResponse(entities *entities.Todo) *fiber.Map {
	return &fiber.Map{
		"title":     entities.Title,
		"is_active": entities.Is_active,
		"priority":  entities.Priority,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func TodoErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  "Failed",
		"message": "err.Error()",
		"data":    "",
	}
}
