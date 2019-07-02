package server

import (
	"mall/app/api/web/user/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var param model.RegisterParam
	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "注册失败,"+err.Error(), nil)
		return
	}

	err := svr.Register(param)
	if err != nil {
		http.Response(c, 400, "注册失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "注册成功", nil)
	return
}
