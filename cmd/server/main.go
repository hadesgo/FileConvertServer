package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/common"
	"github.com/hadesgo/FileConvertServer/router"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = router.CollectRoute(r)

	// 监听端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
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
