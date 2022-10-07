package model3ds

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

    routes := app.Group("/model3ds")
    routes.Post("/", h.AddModel3D)
    routes.Get("/", h.GetModel3Ds)
    routes.Get("/:user", h.GetUserModel3Ds)
    routes.Get("/:id", h.GetModel3D)
    routes.Put("/:id", h.UpdateModel3D)
    routes.Delete("/:id", h.DeleteModel3D)
}