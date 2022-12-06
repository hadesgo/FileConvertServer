package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/common"
	"github.com/hadesgo/FileConvertServer/dto"
	"github.com/hadesgo/FileConvertServer/models"
	"github.com/hadesgo/FileConvertServer/response"
	"github.com/hadesgo/FileConvertServer/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	var ginBindUser = models.User{}
	ctx.Bind(&ginBindUser)

	name := ginBindUser.Name
	email := ginBindUser.Email
	password := ginBindUser.Password

	if !utils.VerifyEmailFormat(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱格式不正确")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, email, password)

	// 验证用户是否已存在
	if isEmailExist(DB, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密失败")
	}

	// 创建用户
	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)

	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var ginBindUser = models.User{}
	ctx.Bind(&ginBindUser)

	email := ginBindUser.Email
	password := ginBindUser.Password

	// 数据验证
	if !utils.VerifyEmailFormat(email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱格式不正确")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 判断用户是否存在
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, errToken := common.ReleaseToken(user)
	if errToken != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "生成token失败")
		log.Printf("token generate error: %v", errToken)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "ok")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(models.User))}, "ok")
}

func isEmailExist(db *gorm.DB, email string) bool {
	var user models.User
	if res := db.Where("email = ?", email).First(&user); res.Error != nil {
		return false
	}

	return true
}
