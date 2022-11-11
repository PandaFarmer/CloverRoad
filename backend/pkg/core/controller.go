package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

)

type handler struct {
	DB *gorm.DB
}

func (h handler) RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	auth_routes := r.Group("/auth/login")
	{
		auth_routes.POST("/", h.Login)
	}
}
