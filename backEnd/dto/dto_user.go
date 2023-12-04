package dto

type UserDto struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Number   int    `json:"number"`
	Name     string `json:"name"`
}

type UsersDto []UserDto
