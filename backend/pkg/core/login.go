package core

import (
	"fmt"
	"time"

	"JWT-auth/internal/middleware"
	"JWT-auth/internal/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/PandaFarmer/CloverRoad/backend-gin/pkg/common/models"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) Login(c *gin.Context) {
	// c.Append("Access-Control-Allow-Origin", "http://localhost:3000/")
	//body := LoginRequestBody{}
	//email := body.Email
	//pass := body.Password
	loginRequestBody := new(LoginRequestBody)

	if err := c.BodyParser(loginRequestBody); err != nil {
		return err
	}
	fmt.Println(loginRequestBody.Email)
	fmt.Println(loginRequestBody.Password)

	email := loginRequestBody.Email
	pass := loginRequestBody.Password

	// email := c.FormValue("email")
	// pass := c.FormValue("pass")

	/*email := c.FormValue("email")
	fmt.Println("email:%s", email)
	pass := c.FormValue("pass")*/

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
	if result := h.DB.First(&user, models.User{Email: email}); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if !ComparePasswords(user.Password, []byte(pass)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	//*/

	// Create the Claims
	/*claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}*/
	claims := jwt.MapClaims{
		"name":  user.UserName,
		"admin": user.IsSuperUser,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Printf("user.UserName: %v/n", user.UserName)
	fmt.Printf("Token: %v/n", token)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// fmt.Printf()
	return c.JSON(fiber.Map{"token": t})
}

func (h handler) Accessible(c *gin.Context) {
	return c.SendString("Accessible")
}

func (h handler) Restricted(c *gin.Context) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(fiber.Map{"Welcome": name})
	// return c.SendString("Welcome " + name)
}
