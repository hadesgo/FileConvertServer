package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/common"
	"github.com/hadesgo/FileConvertServer/controller"
	"github.com/hadesgo/FileConvertServer/middleware"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.Run()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
