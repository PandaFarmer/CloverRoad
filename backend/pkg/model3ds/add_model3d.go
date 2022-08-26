package model3ds

import (
	"github.com/PandaFarmer/CloverRoad/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddModel3DRequestBody struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (h handler) AddModel3D(c *fiber.Ctx) error {
	body := AddModel3DRequestBody{}

	// parse body, attach to AddModel3DRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var model3d models.Model3D

	model3d.Title = body.Title
	model3d.Author = body.Author
	model3d.Description = body.Description
	model3d.Price = body.Price

	// insert new db entry
	if result := h.DB.Create(&model3d); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&model3d)
}
