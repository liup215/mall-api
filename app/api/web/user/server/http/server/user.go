package server

import (
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
)

func info(c *gin.Context) {
	user, _ := c.Get("userID")
	http.Response(c, 200, "成功", gin.H{
		"user": user,
	})
}
