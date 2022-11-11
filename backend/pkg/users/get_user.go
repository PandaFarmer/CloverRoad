package users

import (
	"fmt"
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	fmt.Println("getting user by id")
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &user)
	return
}

func (h handler) GetUserByEmail(c *gin.Context) {
	fmt.Println("getting user by email")
	email := c.Param("email")

	var user models.User

	if result := h.DB.First(&user, email); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.Status(http.StatusOK)
}
