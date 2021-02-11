package app

import "github.com/gin-gonic/gin"

var (
	// router will be in charge of creating different goRoutines for different request we want to handle .
	router = gin.Default() // a private var for app package
)

func StartApplication() {
	mapUrls()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
