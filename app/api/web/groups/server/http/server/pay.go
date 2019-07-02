package server

import (
	"mall/app/api/web/groups/model"
	"mall/lib/net/http"
	"mall/lib/net/http/middleware/auth"

	"github.com/gin-gonic/gin"
)

func prePay(c *gin.Context) {
	u, exist := auth.GetCurrentUser(c)
	if !exist {
		http.Response(c, 401, "用户未登录", nil)
		return
	}

	if u.Usertype == 0 {
		http.Response(c, 400, "暂无权限", nil)
		return
	}

	var p model.PrePayParams
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	token, err := svr.Prepay(u.Openid, p, c.ClientIP())
	if err != nil {
		http.Response(c, 400, "下单错误"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "获取成功", gin.H{
		"token": token,
	})

}
