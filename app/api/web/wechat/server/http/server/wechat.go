package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall/lib/net/http"
	"strconv"
)

func wechatOauthParam(c *gin.Context) {
	uniacidStr := c.DefaultQuery("uniacid", "1")
	uniacid, err := strconv.Atoi(uniacidStr)

	if err != nil {
		http.Response(c, 400, "参数错误", nil)
		return
	}

	module := c.DefaultQuery("module", "")

	r, err := svr.OauthParams(uniacid)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	session := sessions.Default(c)
	session.Set("SSESSION_ID", r.State)
	session.Set("MODULE", module)
	session.Save()

	http.Response(c, 200, "获取成功", r)
}
