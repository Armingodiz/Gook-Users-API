package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// provide endpoints to interact with users API
// every request should be handled by the controller , so this is our entry point of our application

// handle every request for creating user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing ... ")
}

// handle every request for getting user from db
func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implementing ... ")
}

/*func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing ... ")
}*/
