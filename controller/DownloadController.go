package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hadesgo/FileConvertServer/response"
)

func DownloadFile(ctx *gin.Context) {
	zipPath, _ := filepath.Abs("./test/iTopPDF-OCRModule.zip")
	fmt.Println(zipPath)
	//打开文件
	_, errByOpenFile := os.Open(zipPath)
	//非空处理
	if errByOpenFile != nil {
		response.Response(ctx, http.StatusFound, 404, nil, "文件不存在")
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(zipPath)
}
