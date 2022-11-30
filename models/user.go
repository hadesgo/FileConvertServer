package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"type:varchar(256);not null"`
	Password string `gorm:"type:varchar(256);not null"`
}