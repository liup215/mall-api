package server

import (
	"mall/app/api/web/groups/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func getShareVistorInfo(c *gin.Context) {
	var q model.VisitorQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	v, err := svr.ShareVisitorInfo(q)
	if err != nil {
		http.Response(c, 400, "获取失败，"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "获取成功", v)

}
