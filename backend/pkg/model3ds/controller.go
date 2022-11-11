package model3ds

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/model3ds")
	routes.POST("/", h.AddModel3D)
	routes.GET("/", h.GetModel3Ds)
	routes.GET("/:id", h.GetModel3D)
	routes.PUT("/:id", h.UpdateModel3D)
	routes.DELETE("/:id", h.DeleteModel3D)
}
