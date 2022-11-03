package users

import (
	"fmt"
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
)

func (h handler) GetUser(c *gin.Context) {
	fmt.Println("getting user by id")
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

func (h handler) GetUserByEmail(c *gin.Context) {
	fmt.Println("getting user by email")
	email := c.Params("email")

	var user models.User

	if result := h.DB.First(&user, email); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}
