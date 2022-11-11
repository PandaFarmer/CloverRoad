package model3ds

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddModel3DRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddModel3D(c *gin.Context) {
	body := AddModel3DRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var model3d models.Model3D

	model3d.Title = body.Title
	model3d.Author = body.Author
	model3d.Description = body.Description

	if result := h.DB.Create(&model3d); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &model3d)
}
