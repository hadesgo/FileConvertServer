package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" form:"name"`
	Email string `gorm:"type:varchar(256);not null" form:"email"`
	Password string `gorm:"type:varchar(256);not null" form:"password"`
}