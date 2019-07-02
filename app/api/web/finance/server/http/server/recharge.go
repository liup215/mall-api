package server

import (
	"mall/app/api/web/finance/model"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func prerecharge(c *gin.Context) {
	var param model.PreRechargeParam

	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return

	}

	res, err := svr.PreRecharge(param)

	if err != nil {
		http.Response(c, 400, "申请失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "申请成功", res)
	return

}
