package users

import "encoding/json"

type PublicUser struct {
	Id       int64  `json:"id"`
	FirsName string `json:"first_name"`
	Status   string `json:"status"`
}
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirsName    string `json:"first_name"`
	LastNAme    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type Users []User

func (users Users) Marshal(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPublic)
	}
	return result
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		var data PublicUser
		userJson, _ := json.Marshal(user)
		json.Unmarshal(userJson, &data)
		return data
	} else {
		var data PrivateUser
		userJson, _ := json.Marshal(user)
		json.Unmarshal(userJson, &data)
		return data
	}

}
