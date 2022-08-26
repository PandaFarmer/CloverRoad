package main

import (
    "log"
    //"fmt"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/cors"
    //"github.com/gofiber/fiber/v2/middleware/csrf"
    jwtware "github.com/gofiber/jwt/v3"
    "github.com/PandaFarmer/CloverRoad/pkg/model3ds"
    "github.com/PandaFarmer/CloverRoad/pkg/users"
    "github.com/PandaFarmer/CloverRoad/pkg/core"
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
    
    //Logging:
    //app.Use(logger.New())
    

    //app.Use(logger.New(logger.Config{
    //    Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
    //}))

    app.Use(logger.New(logger.Config{
        // For more options, see the Config section
      Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
    }))

    //Cors:
    app.Use(cors.New())
    //app.Use(csrf.New())
    
    core.RegisterRoutes(app, db)
    
    app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

    model3ds.RegisterRoutes(app, db)
    users.RegisterRoutes(app, db)

    app.Listen(c.Port)
}