package dto

import "github.com/hadesgo/FileConvertServer/models"

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserDto(user models.User) UserDto {
	return UserDto{Name: user.Name, Email: user.Email}
}
