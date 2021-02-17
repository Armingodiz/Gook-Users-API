package users

import (
	"github.com/ArminGodiz/Gook-Users-API/datasources/mysql/users_db"
	"github.com/ArminGodiz/Gook-Users-API/logger"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
	"time"
)

// user data access object

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFoundUserByStatus      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying get user from db ", err)
		// *******  we dont pass the err message to user because of security aspects !! ********
		return errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	result := stm.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirsName, &user.LastNAme, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while trying get user from db ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error trying to save user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	user.DateCreated = time.Now().String()
	insertResult, err := stm.Exec(user.FirsName, user.LastNAme, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error trying to save user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error trying to save user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error trying to Update user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	_, err = stm.Exec(user.FirsName, user.LastNAme, user.Email, user.Id)
	if err != nil {
		logger.Error("error trying to Update user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error trying to delete user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	if _, err := stm.Exec(user.Id); err != nil {
		logger.Error("error trying to delete user : ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stm, err := users_db.Client.Prepare(queryFoundUserByStatus)
	if err != nil {
		logger.Error("error trying to find  user by status  : ", err)
		return nil, errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	rows, err := stm.Query(status)
	if err != nil {
		logger.Error("error trying to find  user by status  : ", err)
		return nil, errors.NewInternalServerError("error in DB !")
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		var newUser User
		if err := rows.Scan(&newUser.Id, &newUser.FirsName, &newUser.LastNAme, &newUser.Email, &newUser.DateCreated, &newUser.Status); err != nil {
			logger.Error("error trying to find  user by status  : ", err)
			return nil, errors.NewInternalServerError("error in DB !")
		}
		result = append(result, newUser)
	}
	if len(result) == 0 {
		return nil, errors.NewInternalServerError("no match while trying to find  user by status  ")
	}
	return result, nil
}
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stm, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error while trying get user from db ", err)
		// *******  we dont pass the err message to user because of security aspects !! ********
		return errors.NewInternalServerError("error in DB !")
	}
	defer stm.Close()
	result := stm.QueryRow(user.Email, user.Password)
	if err := result.Scan(&user.Id, &user.FirsName, &user.LastNAme, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while trying get user from db ", err)
		return errors.NewInternalServerError("error in DB !")
	}
	return nil
}
