package core

import (
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

    routes := app.Group("/login")
    routes.Get("/:email.:pass", h.Login)
}