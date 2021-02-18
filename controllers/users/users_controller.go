package users

import (
	"github.com/ArminGodiz/Gook-Users-API/domain/users"
	"github.com/ArminGodiz/Gook-Users-API/services"
	"github.com/ArminGodiz/Gook-Users-API/services/oauth"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// provide endpoints to interact with users API
// every request should be handled by the controller , so this is our entry point of our application

// handle every request for creating user

func getUserID(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number ")
	}
	return userID, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body !")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	result, saveErr := services.UsersService.CreatUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	//fmt.Println(result)
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

// handle every request for getting user from db
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Code, err)
		return
	}
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
	}
	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(false))
	}
	c.JSON(http.StatusOK, user.Marshal(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body !")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	user.Id = userID

	isPArtial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPArtial, user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Code, err)
		return
	}
	// dont use c.String because we should use the same type of body answer for an endpoint
	c.JSON(http.StatusOK, map[string]string{"status": "deleted !"})
}
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var req users.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		err2 := errors.NewBadRequestError("INVALID JSON BODY !")
		c.JSON(err2.Code, err2)
		return
	}
	user, err := services.UsersService.LoginUser(req)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}
