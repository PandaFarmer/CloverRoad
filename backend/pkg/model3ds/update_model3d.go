package model3ds

import (
	"github.com/PandaFarmer/CloverRoad/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateModel3DRequestBody struct {
	Title                string  `json:"title"`
	Author               string  `json:"author"`
	Description          string  `json:"description"`
	Price                float64 `json:"price"`
	SerializedFile3D     []byte  `json:"serialized_file_3d"`
	FileNameAndExtension string  `json:"file_name_and_extension"`
}

func (h handler) UpdateModel3D(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateModel3DRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var model3d models.Model3D

	if result := h.DB.First(&model3d, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	model3d.Title = body.Title
	model3d.Author = body.Author
	model3d.Description = body.Description
	model3d.Price = body.Price
	model3d.SerializedFile3D = body.SerializedFile3D
	model3d.FileNameAndExtension = body.FileNameAndExtension

	// save model3d
	h.DB.Save(&model3d)

	return c.Status(fiber.StatusOK).JSON(&model3d)
}
