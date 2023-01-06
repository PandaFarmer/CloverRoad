package model3ds

import (
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/config"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TestRegisterRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
		db  *gorm.DB
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		requestType  string
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/model3ds",
			expectedCode: 200,
			requestType:  "GET",
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route does not exist",
			route:        "/not-found",
			expectedCode: 404,
			requestType:  "GET",
		},
	}

	os.Chdir("../../")
	c, err := config.LoadConfig()
	// Define Fiber app.
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db, err := db.Init(c.DBUrl)
	if err != nil {
		log.Fatalln("Failed at db.Init", err)
	}
	RegisterRoutes(app, db)

	for _, tt := range tests {
		bodyReader := strings.NewReader(`{"Username": "12124", "Password": "testinasg", "Channel": "M"}`)
		req := httptest.NewRequest(tt.requestType, tt.route, bodyReader)
		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.description)
		// t.Run(tt.name, func(t *testing.T) {
		// RegisterRoutes(tt.args.app, tt.args.db)

		// })
	}
}
