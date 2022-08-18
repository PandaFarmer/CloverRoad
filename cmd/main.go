package main

import (
    "log"
    //"fmt"

    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/model3ds"
    "github.com/PandaFarmer/CloverRoad/pkg/users"
    "github.com/PandaFarmer/CloverRoad/pkg/common/config"
    "github.com/PandaFarmer/CloverRoad/pkg/common/db"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    app := fiber.New()
    db := db.Init(c.DBUrl)

    //app.Get("/", func(ctx *fiber.Ctx) error {
    //    return ctx.Status(fiber.StatusOK).SendString(c.Port)
    //})

    books.RegisterRoutes(app, db)
    users.RegisterRoutes(app, db)

    app.Listen(c.Port)
}