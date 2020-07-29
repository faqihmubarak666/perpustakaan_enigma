package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AllUser struct {
	IdUser   string `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
}
