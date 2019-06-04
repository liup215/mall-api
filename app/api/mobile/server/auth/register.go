package auth

import (
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
)

func Register(c *gin.Context) {
	type Register struct {
		Uniacid    int    `json:"uniacid" form:"uniacid" binding:"required"`
		Mobile     string `json:"mobile" form:"mobile" binding:"required"`
		Pwd        string `json:"pwd" form:"pwd" binding:"required"`
		Pwdconfirm string `json:"pwdconfirm" form:"pwdconfirm" binding "required"`
	}

	var r Register
	if err := c.Bind(&r); err != nil {
		http.Response(c, 400, "参数错误，"+err.Error(), nil)
		return
	}

	if r.Pwd != r.Pwdconfirm {
		http.Response(c, 400, "两次输入密码不一致", nil)
	}

	err := authSvr.Register(r.Uniacid, r.Mobile, r.Pwd)
	if err != nil {
		http.Response(c, 400, "注册失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "注册成功")

}
