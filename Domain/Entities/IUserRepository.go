package Entities

import (
	dto "proyecto.com/Domain/Dto"
	"proyecto.com/Infraestructure/Database"
)

type IUserRepository interface {
	GetUser(user *dto.UserDto) (User, error)
	SignUp(user *User) bool
}

// implementar interfaz
func (u User) GetUser(user *dto.UserDto) (User, error) {
	var userStored User
	result := Database.DB.Where("email = ?", user.Email).First(&userStored)

	return userStored, result.Error
}

func (u User) SignUp(user *User) bool {

	var userStored User

	result := Database.DB.Where("email = ?", user.Email).First(userStored)

	if result.Error == nil {
		return false
	}

	createdUser := Database.DB.Create(user)

	err := createdUser.Error

	if err != nil {
		return false
	}
	return true
}
