package handlers

import (
	"fmt"
	"go_todo/api/presenter"
	"go_todo/pkg/activity"
	"go_todo/pkg/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AddBook is handler/controller which creates Books in the BookShop
func AddActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(entities.Activity)
		errbody := c.BodyParser(&requestBody)
		if errbody != nil {
			c.Status(http.StatusBadRequest)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ActivityErrorResponse(errbody))
		}
		requestBody.ID = uuid.New()
		result, err := service.InsertActivity(requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ActivityErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.ActivitySuccessResponse(result))
	}
}
func UpdateActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Activity
		id := c.Params("id")

		errrbody := c.BodyParser(&requestBody)
		if errrbody != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ActivityErrorResponse(errrbody))
		}
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		requestBody.ID = idparam

		result, err := service.UpdateActivity(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		return c.Status(fiber.StatusOK).JSON(presenter.ActivitySuccessResponse(result))
	}
}

// UpdateBook is handler/controller which updates data of Books in the BookShop

// RemoveBook is handler/controller which removes Books from the BookShop
func RemoveActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		errservice := service.RemoveActivity(idparam)
		if errservice != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "Success",
			"message": fmt.Sprintf("Success Deleted %v", id),
		})
	}
}

// GetBooks is handler/controller which lists all Books from the BookShop
func GetActivity(service activity.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		idparam, erruuid := uuid.Parse(id)
		if erruuid != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		activity, errservice := service.GetActivity(idparam)
		if errservice != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %v Not Found", id),
			})
		}
		return c.JSON(presenter.ActivitySuccessResponse(activity))
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
