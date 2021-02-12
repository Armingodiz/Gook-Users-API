package app

import "github.com/ArminGodiz/Gook-Users-API/controllers/ping"
import "github.com/ArminGodiz/Gook-Users-API/controllers/users"

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// user controllers :
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	//router.GET("/users/search", controllers.SearchUser)
}
