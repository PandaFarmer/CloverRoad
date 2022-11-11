package users

import (
	"fmt"
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
	fmt.Println("getting users")
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.Status(http.StatusOK)
}
