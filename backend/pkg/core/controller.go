package core

import (
	"github.com/gin-gonic/gin"
	"JWT-auth/internal/middleware"
	"JWT-auth/internal/routes"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	// routes := app.Group("/login", jwtware.New(jwtware.Config{
	//     KeyFunc : func (ctx *fiber.Ctx) error {}
	// }))
	routes := app.Group("/auth/login")
	routes.Post("/", h.Login)
	routes = app.Group("/auth/restricted", jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	routes.Get("/", h.Restricted)
}
