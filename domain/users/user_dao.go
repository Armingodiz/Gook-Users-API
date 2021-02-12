package users

import (
	"fmt"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"time"
)

// user data access object

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("email %s already registred ", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprint("user %d not already request", user.Id))
	}
	now := time.Now()
	user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprint("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirsName = result.FirsName
	user.LastNAme = result.LastNAme
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
