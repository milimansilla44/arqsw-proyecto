package user

import (
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)
	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("Users: ", users)
	return users
}

func GetUserByUserName(LoginDto dto.LoginDto) model.User {
	var user model.User
	Db.First(&user, "UserName = ? AND password = ?", LoginDto.UserName, LoginDto.Password)
	log.Debug("User: ", user)
	return user
}

func NewUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("ERROR al crear el usuario")
		return user
	}
	log.Debug("User Created: ", user.Id)
	return user
}
