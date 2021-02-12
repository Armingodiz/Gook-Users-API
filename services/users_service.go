package services

import (
	"github.com/ArminGodiz/Gook-Users-API/domain/users"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
)

func CreatUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
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

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	current.FirsName = user.FirsName
	current.LastNAme = user.LastNAme
	current.Email = user.Email
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
