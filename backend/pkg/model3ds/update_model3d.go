package model3ds

import (
    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/common/models"
)

type UpdateModel3DRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price 		float64  `json:"price"`
    BlobData    []byte  `json:"blob_data"`
    FileName    string  `json:"file_name"`
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
    model3d.BlobData = body.BlobData
    model3d.FileName = body.FileName

    // save model3d
    h.DB.Save(&model3d)

    return c.Status(fiber.StatusOK).JSON(&model3d)
}