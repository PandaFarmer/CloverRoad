package model3ds

import (
	"log"
	"testing"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/config"
	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Test_handler_AddModel3D(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		requestType  string
		name         string
		fields       fields
		args         args
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			description:  "get HTTP status 200 from post model3ds route",
			route:        "/model3ds",
			expectedCode: 200,
			requestType:  "POST",
			name:         "posting model3ds",
			fields:       fields{},
			args:         args{},
			wantErr:      false,
		},
	}

	c, err := config.LoadConfig()
	// Define Fiber app.
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	RegisterRoutes(app, db)

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			h := handler{
				DB: tt.fields.DB,
			}
			if err := h.AddModel3D(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("handler.AddModel3D() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
