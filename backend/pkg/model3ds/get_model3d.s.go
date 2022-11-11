package model3ds

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetModel3Ds(c *gin.Context) {
	var model3ds []models.Model3D

	if result := h.DB.Find(&model3ds); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &model3ds)
}
