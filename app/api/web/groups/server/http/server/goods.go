package server

import (
	"mall/app/api/web/groups/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func goodsQueryList(c *gin.Context) {
	var q model.GoodsQuery
	if err := c.Bind(&q); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	list, total := svr.QueryGoodsList(q)
	http.Response(c, 200, "获取成功", gin.H{
		"list":  list,
		"total": total,
	})
}
