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

func userCheckWechat(c *gin.Context) {
	var param model.UserCheckWechatParam
	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "用户校验失败,"+err.Error(), nil)
		return
	}

	member, err := svr.UserCheckWechat(param)
	if err != nil {
		http.Response(c, 400, "用户校验失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "用户校验成功", member)
}

func memberUpdatePwd(c *gin.Context) {
	var p model.UpdatePwdParam
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "更新失败,"+err.Error(), nil)
		return
	}

	err := svr.MemberUpdatePwd(&p)
	if err != nil {
		http.Response(c, 400, "更新失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "更新成功", nil)

}
