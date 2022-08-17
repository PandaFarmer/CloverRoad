package model3ds

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/common/models"
)

func (h handler) GetModel3Ds(c *fiber.Ctx) error {
    var model3ds []models.Model3D

    if result := h.DB.Find(&model3ds); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusOK).JSON(&model3ds)
}