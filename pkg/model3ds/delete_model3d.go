package model3ds

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/common/models"
)

func (h handler) DeleteModel3D(c *fiber.Ctx) error {
    id := c.Params("id")

    var model3d models.Model3D

    if result := h.DB.First(&model3d, id); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    // delete model3d from db
    h.DB.Delete(&model3d)

    return c.SendStatus(fiber.StatusOK)
}