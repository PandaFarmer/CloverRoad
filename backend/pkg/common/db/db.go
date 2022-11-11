package db

import (
	"log"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var GlobalDB *gorm.DB

func Init(dbUrl string) (err error) {

	dsn := "host=localhost user=postgres password=3t4t5LQS dbname=clover port=5432 sslmode=disable"
	GlobalDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	GlobalDB.AutoMigrate(&models.Model3D{})
	GlobalDB.AutoMigrate(&models.User{})

	return
}
