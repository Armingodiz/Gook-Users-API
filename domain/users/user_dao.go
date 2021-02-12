package users

import (
	"fmt"
	"github.com/ArminGodiz/Gook-Users-API/datasources/mysql/users_db"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"time"
)

// user data access object

const (
	queryInsertUser = ("INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);")
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stm.Close()
	user.DateCreated = time.Now().String()
	insertResult, err := stm.Exec(user.FirsName, user.LastNAme, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error trying to save user ")

	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
