package users

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/core"
    "github.com/PandaFarmer/CloverRoad/pkg/common/models"
)

type AddUserRequestBody struct {
    Id          int    `json:"id" gorm:"primaryKey"`
	Email		string `json:"email"`
	UserName	string `json:"user_name"`
	FirstName   string `json:"first_name"`
	Surname	string `json:"surname"`
	Password	string `json:"password"`
	DateJoined  time.Time   `json:"date_joined"`
	IsSuperUser	bool 	`json:"is_super_user"`
}

func (h handler) AddUser(c *fiber.Ctx) error {
    body := AddUserRequestBody{}

    // parse body, attach to AddUserRequestBody struct
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    var user models.User

    user.Id = body.Id
    user.Email = body.Email
    user.UserName = body.UserName
    user.FirstName = body.FirstName
    user.Surname = body.Surname
    user.Password = core.HashAndSalt([]byte(body.Password))
    user.DateJoined = time.Now()
    user.IsSuperUser = body.IsSuperUser

    // insert new db entry
    if result := h.DB.Create(&user); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(&user)
}