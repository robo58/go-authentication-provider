package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robo58/go-authentication-provider/config"
	"github.com/robo58/go-authentication-provider/data"
	"github.com/robo58/go-authentication-provider/router"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	data.SetupDB()
	data.SetupHydra()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func main() {
	configPath := "./config.yml"
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
