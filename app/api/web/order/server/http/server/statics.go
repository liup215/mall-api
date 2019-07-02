package server

import (
	"mall/app/api/web/order/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func orderStatics(c *gin.Context) {
	var p model.OrderStaticsParam
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "参数错误, "+err.Error(), nil)
		return
	}

	r, err := svr.OrderStatics(p)

	if err != nil {
		http.Response(c, 400, "获取失败, "+err.Error(), nil)
		return
	}
	http.Response(c, 200, "获取成功", r)
}
