package app

import "github.com/ArminGodiz/Gook-Users-API/controllers/ping"
import "github.com/ArminGodiz/Gook-Users-API/controllers/users"

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// user controllers :
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	//router.GET("/users/search", controllers.SearchUser)
}
