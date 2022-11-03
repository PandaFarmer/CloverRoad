package main

import (
	"github.com/PandaFarmer/CloverRoad/backend/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/model3ds"
	"github.com/PandaFarmer/CloverRoad/backend/pkg/users"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	users.RegisterRoutes(r, h)
	model3ds.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
