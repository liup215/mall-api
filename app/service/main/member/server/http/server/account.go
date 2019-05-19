package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall/app/service/main/member/model"
	"mall/lib/net/http"
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

func login(c *gin.Context) {
	var param model.LoginParam
	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "登录失败,"+err.Error(), nil)
		return
	}

	cookie, err := svr.Login(param)
	if err != nil {
		http.Response(c, 400, "登录失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "登录成功", gin.H{
		fmt.Sprintf("be50___ewei_shopv2_member_session_%v", param.Uniacid): cookie,
	})
	return
}
