package users

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend-gin/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteUser(c *gin.Context) {
	id := c.Params("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// delete user from db
	h.DB.Delete(&user)

	return c.SendStatus(fiber.StatusOK)
}
