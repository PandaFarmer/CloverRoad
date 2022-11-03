package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	fmt.Println("Initialized users endpoint...")
	routes := app.Group("/users")
	routes.Post("/", h.AddUser)
	routes.Get("/", h.GetUsers)
	routes.Get("/:id", h.GetUser)
	routes.Get("/:email", h.GetUserByEmail)
	routes.Put("/:id", h.UpdateUser)
	routes.Delete("/:id", h.DeleteUser)
}
