package users

import (
	"fmt"

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
	fmt.Println("Initialized users endpoint...")
	users_routes := r.Group("/users")
	{
		users_routes.POST("/", h.AddUser)
		users_routes.GET("/", h.GetUsers)
		users_routes.GET("/by-id/:id", h.GetUser)
		users_routes.GET("/by-email/:email", h.GetUserByEmail)
		users_routes.PUT("/:id", h.UpdateUser)
		users_routes.DELETE("/:id", h.DeleteUser)
	}
}
