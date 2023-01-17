package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/config"
	"tiktok/controller"
	"tiktok/log"
	"tiktok/models"
)

func main() {
	log.Init()
	config.Init()
	models.Init()

	e := gin.Default()
	controller.InitRouter(e)

	//todo: run message websocket server here

	err := e.Run(":" + config.Conf.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
