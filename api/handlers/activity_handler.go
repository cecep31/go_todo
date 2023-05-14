package handlers

import (
	"fmt"
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
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ActivityErrorResponse(err))
		}
		result, err := service.InsertActivity(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.ActivitySuccessResponse(result))
	}
}
func UpdateActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Activity
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenter.ActivityErrorResponse(err))
		}

		errrbody := c.BodyParser(&requestBody)

		if errrbody != nil {
			c.Status(http.StatusBadRequest)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ActivityErrorResponse(errrbody))
		}

		requestBody.ID = uint(id)

		result, err := service.UpdateActivity(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ActivitySuccessResponse(result))
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
		errservice := service.RemoveActivity(uint(id))
		if errservice != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"status":  "Failed",
				"message": fmt.Sprintf("Failed Deleted %v", id),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "Deleted",
			"message": fmt.Sprintf("Success Deleted %v", id),
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenter.ActivityErrorResponse(err))
		}
		activity, errservice := service.GetActivity(uint(id))
		if errservice != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenter.ActivityErrorResponse(errservice))
		}
		return c.JSON(activity)
	}
}

func GetActivities(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activity, err := service.GetActivities()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ActivityErrorResponse(err))
		}
		return c.JSON(presenter.ActivitiesSuccessResponse(activity))
	}
}
