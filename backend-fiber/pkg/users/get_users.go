package users

import (
	"fmt"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUsers(c *fiber.Ctx) error {
	fmt.Println("getting users")
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&users)
}
