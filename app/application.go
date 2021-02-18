package app

import (
	"github.com/ArminGodiz/Gook-Users-API/logger"
	"github.com/gin-gonic/gin"
)

var (
	// router will be in charge of creating different goRoutines for different request we want to handle .
	router = gin.Default() // a private var for app package
)

func StartApplication() {
	mapUrls()
	// for cloud to hit it and understand our web server is still running
	logger.Info("starting User API service ...")
	err := router.Run(":1111")
	if err != nil {
		panic(err)
	}
}
