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
	var user Userinfo
	var userinfos []Userinfo
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr(err)

	// 添加数据
	createTime, _ := time.Parse("2006-01-02", "2019-12-12")
	user = Userinfo{Username: "lys", Department: "RD", Created: createTime}
	result := db.Create(&user)
	if result.Error == nil {
		fmt.Println(user.Uid)
		fmt.Println(result.RowsAffected)
	}

	// 更新数据
	db.First(&user)
	user.Username = "lysup"
	db.Save(&user)

	// 删除id是2的行
	user = Userinfo{Uid: 1}
	db.Delete(&user)

	// 查询数据
	result = db.Find(&userinfos)
	if result.Error == nil {
		fmt.Println(result.RowsAffected)
		for _, v := range userinfos {
			fmt.Println(v.Uid)
			fmt.Println(v.Username)
			fmt.Println(v.Department)
			fmt.Println(v.Created)
		}
	}
}
