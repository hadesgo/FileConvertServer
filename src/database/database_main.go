package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Userinfo struct {
	Uid        uint `gorm:"primaryKey"`
	Username   string
	Department string
	Created    time.Time
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr(err)

	createTime, _ := time.Parse("2006-01-02", "2019-12-12")
	user := Userinfo{Username: "lys", Department: "RD", Created: createTime}
	result := db.Create(&user)
	fmt.Println(user.Uid)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}
