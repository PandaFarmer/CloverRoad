package main

import (
	// "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/config"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/db"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/core"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/model3ds"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/users"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	// jwtware "github.com/gofiber/jwt/v3"
)

// const jwtSecret = "asecret"

// func authRequired() func(ctx *fiber.Ctx) error {

// 	return jwtware.New(jwtware.Config{
// 		// ErrorHandler: func(ctx *fiber.Ctx, err error) {
// 		// 	ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		// 		"error": "Unauthorized",
// 		// 	})
// 		// },
// 		SigningKey: []byte(jwtSecret),
// 	})
// }

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	//	use form to receive filedata instead of json, this ReadBufferSize change is not need 4now
	// app := fiber.New(fiber.Config{
	// 	Prefork:       true,
	// 	CaseSensitive: true,
	// 	StrictRouting: true,
	// 	ReadBufferSize: 32*1024*1024,
	// 	ServerHeader:  "Fiber",
	// 	AppName: "CloverRoad v0.01",
	// })
	db := db.Init(c.DBUrl)

	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "pid: ${pid}${locals:requestid} ${status} - ${method} ${path}" +
			"\nreqHeaders:${reqHeaders}\nqueryParams:${queryParams}\nbody:${body}" +
			"\nquery:{query:}\nform:{form:}\ncookie:{cookie:}\nlocals:{locals:}" +
			"\nrespHeader:${respHeader:}" +
			"\nua:${ua}\nerror:{error}\n",
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	// 	AllowHeaders:     "*", //"Origin, Content-Type, Accept, Authorization, other_header",
	// 	AllowCredentials: true,
	// }))

	core.RegisterRoutes(app, db)

	// app.Use(csrf.New())

	SecureEndpointInitialization := func() {
		// shit := jwtware.New(jwtware.Config{
		// 	SigningKey: []byte("secret"),
		// 	// TokenLookup: "param:Authorization",
		// })
		// fmt.Println()
		// app.Use(shit)

		model3ds.RegisterRoutes(app, db)
		// fmt.Println("About to register users endpoint")
		users.RegisterRoutes(app, db)

		app.Listen(c.Port)
	}
	defer SecureEndpointInitialization()
}
