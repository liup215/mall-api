package server

import (
	"mall/app/api/web/wechat/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func userInfo(c *gin.Context) {

	var p model.UserInfoParams
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	fans, err := svr.GetUserInfo(p)
	if err != nil {
		http.Response(c, 400, "获取错误,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "获取成功", fans)
}
