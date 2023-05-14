package handlers

import (
	"go_todo/api/presenter"
	"go_todo/pkg/activity"
	"go_todo/pkg/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AddBook is handler/controller which creates Books in the BookShop
func AddActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Activity
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		result, err := service.InsertActivity(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitySuccessResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		err = service.RemoveActivity(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		activity, err := service.GetActivity(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(activity)
	}
}
