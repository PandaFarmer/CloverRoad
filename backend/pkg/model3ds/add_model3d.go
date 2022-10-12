package model3ds

import (
	"fmt"

	"github.com/PandaFarmer/CloverRoad/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddModel3DRequestBody struct {
	Title                string  `form:"title"`
	Author               string  `form:"author"`
	Description          string  `form:"description"`
	Price                float64 `form:"price"`
	SerializedFile3D     []byte  `form:"serialized_file_3d"`
	FileNameAndExtension string  `form:"file_name_and_extension"`
}

func (h handler) AddModel3D(c *fiber.Ctx) error {
	fmt.Println("herro adding model3d")
	body := AddModel3DRequestBody{}

	fmt.Println("something wrong with the received request body?")
	// parse body, attach to AddModel3DRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var model3d models.Model3D

	fmt.Println("hello wtf")
	// fmt.Println(body.SerializedFile3D)
	fmt.Println(body)
	fmt.Println(body.FileNameAndExtension)
	fmt.Println("hello wtf")

	// copyFields := make(map[string][]byte)
	// for k, v := range fields{
	// 	var copyField = make([]byte, len(v))
	// 	copy(copyField, v)
	// 	fmt.Println(string(copyField))
	// 	copyFields[k] = copyField
	// }

	model3d.Title = body.Title
	model3d.Author = body.Author
	model3d.Description = body.Description
	model3d.Price = body.Price
	model3d.SerializedFile3D = body.SerializedFile3D         //sus
	model3d.FileNameAndExtension = body.FileNameAndExtension //sus

	fmt.Println(model3d.Title)

	// insert new db entry
	if result := h.DB.Create(&model3d); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&model3d)
}
