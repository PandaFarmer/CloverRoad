package main

import (
    //"log"
    "fmt"

    "github.com/gofiber/fiber/v2"
    "github.com/PandaFarmer/CloverRoad/pkg/model3ds"
    "github.com/PandaFarmer/CloverRoad/pkg/common/config"
    "github.com/PandaFarmer/CloverRoad/pkg/common/db"
    //"../pkg/model3ds"
    //"../pkg/common/config"
    //"../pkg/common/db"
)

func main() {
    fmt.Println("hellow")
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    app := fiber.New()
    db := db.Init(c.DBUrl)
    model3ds.RegisterRoutes(app, db)
    app.Listen(c.Port)
}