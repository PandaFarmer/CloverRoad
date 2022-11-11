package model3ds

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateModel3DRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateModel3D(c *gin.Context) {
	id := c.Param("id")
	body := UpdateModel3DRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var model3d models.Model3D

	if result := h.DB.First(&model3d, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	model3d.Title = body.Title
	model3d.Author = body.Author
	model3d.Description = body.Description

	h.DB.Save(&model3d)

	c.JSON(http.StatusOK, &model3d)
}
