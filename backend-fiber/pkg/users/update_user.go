package users

import (
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/core"
	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequestBody struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	// FirstName   string `json:"first_name"`
	// Surname	string `json:"surname`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	IsSuperUser bool   `json:"is_super_user"`
}

func (h handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateUserRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	user.Id = body.Id
	user.Email = body.Email
	user.UserName = body.UserName
	// user.FirstName = body.FirstName
	// user.Surname = body.Surname
	user.FullName = body.FullName
	user.Password, _ = core.HashAndSalt(body.Password)
	user.IsSuperUser = body.IsSuperUser

	// save user
	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(&user)
}
