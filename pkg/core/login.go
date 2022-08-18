package core

import (
	"time"

	"github.com/gofiber/fiber/v2"
	//jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	"github.com/PandaFarmer/CloverRoad/pkg/common/models"
	
)

func (h handler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("pass")

	/*
	user := c.FormValue("user")
	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	*/

	///*
	//email = c.Params("email")//or just user

    var user models.User
	
    if result := h.DB.First(&user, email); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

	if !ComparePasswords(user.Password, []byte(pass)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	//*/

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}