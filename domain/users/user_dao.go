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
	queryGetUser    = ("SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;")
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("error when getting user from db " + err.Error())
	}
	defer stm.Close()
	result := stm.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirsName, &user.LastNAme, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("error when getting user from db " + err.Error())
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("error trying to save user : " + err.Error())
	}
	defer stm.Close()
	user.DateCreated = time.Now().String()
	insertResult, err := stm.Exec(user.FirsName, user.LastNAme, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError("error trying to save user : " + err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error trying to save user : " + err.Error())

	}
	user.Id = userId
	return nil
}
