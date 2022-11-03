package users

import (
	"fmt"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUser(c *fiber.Ctx) error {
	fmt.Println("getting user by id")
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func (h handler) GetUserByEmail(c *fiber.Ctx) error {
	fmt.Println("getting user by email")
	email := c.Params("email")

	var user models.User

	if result := h.DB.First(&user, email); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
