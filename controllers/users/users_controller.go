package users

import (
	"github.com/ArminGodiz/Gook-Users-API/domain/users"
	"github.com/ArminGodiz/Gook-Users-API/services"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// provide endpoints to interact with users API
// every request should be handled by the controller , so this is our entry point of our application

// handle every request for creating user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body !")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.CreatUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	//fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

// handle every request for getting user from db
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number ")
		c.JSON(err.Code, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number ")
		c.JSON(err.Code, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body !")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user.Id = userID
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

/*func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing ... ")
}*/
