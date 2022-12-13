package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/controller"
	"github.com/hadesgo/FileConvertServer/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.GET("/download/iTopPDF-OCRModule.zip", controller.DownloadFile)

	return r
}
