package services

import (
	"github.com/ArminGodiz/Gook-Users-API/domain/users"
	"github.com/ArminGodiz/Gook-Users-API/utils/errors"
)

func CreatUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
