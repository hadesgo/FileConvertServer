package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/utils"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		email := ctx.PostForm("email")
		password := ctx.PostForm("password")

		// 数据验证
		if !utils.VerifyEmailFormat(email) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "邮箱格式错误!"})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码不能少于6位!"})
			return
		}

		if len(name) == 0 {
			name = utils.RandomString(10)
		}

		log.Println(name, email, password)

		// 创建用户

		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
