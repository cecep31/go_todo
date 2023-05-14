package presenter

import (
	"go_todo/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Book is the presenter object which will be passed in the response by Handler
type Activity struct {
	ID        uuid.UUID `gorm:"primarykey;column:activity_id" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handle
func ActivitySuccessResponse(data *entities.Activity) *fiber.Map {
	activity := Activity{
		ID:        data.ID,
		Title:     data.Title,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return &fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func ActivitiesSuccessResponse(activities *[]entities.Activity) *fiber.Map {
	return &fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activities,
	}
}

func ActivitySuccessCreateResponse(entities *entities.Activity) *fiber.Map {
	return &fiber.Map{
		"title": entities.Title,
		"email": entities.Email,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ActivityErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  "Failed",
		"message": err.Error(),
		"data":    "",
	}
}
