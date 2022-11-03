package model3ds

import (
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUserModel3Ds(c *fiber.Ctx) error {
	user := c.Params("user")

	var model3ds []models.Model3D

	if result := h.DB.Where("author = ?", user).Find(&model3ds); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&model3ds)
}
