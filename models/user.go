package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null" form:"name" json:"name" uri:"name" xml:"name"`
	Email    string `gorm:"type:varchar(256);not null" form:"email" json:"email" uri:"email" xml:"email" binding:"required"`
	Password string `gorm:"type:varchar(256);not null" form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}
