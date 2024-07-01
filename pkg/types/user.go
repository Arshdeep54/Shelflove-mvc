package types

type User struct {
	Id           int32  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsAdmin      bool   `json:"isAdmin"`
	AdminRequest bool   `json:"adminRequest"`
}

type RegisterUserPayload struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}
