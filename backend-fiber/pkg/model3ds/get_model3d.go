package model3ds

import (
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetModel3D(c *fiber.Ctx) error {
	id := c.Params("id")

	var model3d models.Model3D

	if result := h.DB.First(&model3d, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&model3d)
}
