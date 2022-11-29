package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Pool *redis.Pool
)

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) (string, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data string
	
	data, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}

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

func testMysql() {
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

func testRedis() {
	test, err := Get("test")
	fmt.Println(test, err)
}

func main() {
	// testMysql()
	testRedis()
}
