package services

import (
	"github.com/golang-jwt/jwt"
	userCliente "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/user"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"
	crypto "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/encriptado"
	e "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/errors"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError)
	NewUser(usersDto dto.UserDto) (dto.UserDto, e.ApiError)
}

var jwtKey = []byte("secret_key")

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Id = user.Id
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Number = user.Number
	userDto.Street = user.Street
	userDto.City = user.City
	userDto.Name = user.Name

	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.Id
		userDto.UserName = user.UserName
		userDto.Password = user.Password
		userDto.Number = user.Number
		userDto.Street = user.Street
		userDto.City = user.City
		userDto.Name = user.Name
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.Token, e.ApiError) {
	loginDto.Password = crypto.SSHA256(loginDto.Password)
	var user model.User = userCliente.GetUserByUserName(loginDto)

	var tokenDto dto.Token

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("User not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.Id = user.Id
	}
	return tokenDto, nil
}

func (s *userService) NewUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {
	var user model.User

	user.Name = userDto.Name
	user.City = userDto.City
	user.Number = userDto.Number
	user.Password = crypto.SSHA256(userDto.Password)
	user.UserName = userDto.UserName
	user.Street = userDto.Street

	user = userCliente.NewUser(user)

	userDto.Password = user.Password
	userDto.Id = user.Id

	return userDto, nil

}
