package users

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/common/models"
)

func (h handler) GetUser(c *fiber.Ctx) error {
    id := c.Params("id")

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusOK).JSON(&user)
}

func (h handler) GetUserByEmail(c *fiber.Ctx) error {
    email := c.Params("email")

    var user models.User

    if result := h.DB.First(&user, email); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusOK).JSON(&user)
}