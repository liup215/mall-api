package server

import (
	"mall/app/api/web/member/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func creditUpdate(c *gin.Context) {
	var p model.UpdateCreditParam
	if err := c.Bind(&p); err != nil {
		http.Response(c, 400, "操作失败,"+err.Error(), nil)
		return
	}

	err := svr.UpdateCredit(p)

	if err != nil {
		http.Response(c, 400, "操作失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "操作成功", nil)
	return

}
