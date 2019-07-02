package server

import (
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
	"mall/lib/strings"
)

func uploadToken(c *gin.Context) {
	token := svr.UploadToken()
	http.Response(c, 200, "获取成功", gin.H{
		"token": token,
		"key":   strings.CreateNo(""),
	})
}
