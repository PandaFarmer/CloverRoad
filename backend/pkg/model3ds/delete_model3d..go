package model3ds

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteModel3D(c *gin.Context) {
	id := c.Param("id")

	var model3d models.Model3D

	if result := h.DB.First(&model3d, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&model3d)

	c.Status(http.StatusOK)
}
