package users

type PublicUser struct {
	Id          int64  `json:"id"`
	FirsName    string `json:"first_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}
type PrivateUser struct {
	Id          int64  `json:"id"`
	FirsName    string `json:"first_name"`
	LastNAme    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}
