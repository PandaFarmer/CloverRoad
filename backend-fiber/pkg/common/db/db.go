package db

import (
	"log"

	"github.com/PandaFarmer/CloverRoad/backend-fiber/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) (*gorm.DB, error) {
	// dsn := "host=localhost user=postgres password=3t4t5LQS dbname=clover port=5432 sslmode=disable"
	dsn := "host=172.17.0.2 user=postgres password=3t4t5LQS dbname=clover port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	db.AutoMigrate(&models.Model3D{})
	db.AutoMigrate(&models.User{})

	return db, nil
}
