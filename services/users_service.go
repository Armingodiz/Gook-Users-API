package services

import (
	"github.com/ArminGodiz/Gook-Users-API/domain/users"
	"github.com/ArminGodiz/Gook-Users-API/utils/crypto"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
)

func CreatUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userID,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirsName != "" {
			current.FirsName = user.FirsName
		}
		if user.LastNAme != "" {
			current.LastNAme = user.LastNAme
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirsName = user.FirsName
		current.LastNAme = user.LastNAme
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userID int64) *errors.RestErr {
	result := &users.User{
		Id: userID,
	}
	return result.Delete()
}
func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
