package users
// data transfer object == > the object that
//we are going to be transferring from the persistence layer to application backward
import (
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirsName    string `json:"first_name"`
	LastNAme    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address !")
	}
	return nil
}
