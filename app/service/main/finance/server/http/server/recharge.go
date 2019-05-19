package server

import (
	"github.com/gin-gonic/gin"
	"mall/app/service/main/finance/model"
	"mall/lib/net/http"
)

func rechargeSubmit(c *gin.Context) {
	var param model.RechargeParam

	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return

	}

	err := svr.RechargeSubmit(param)

	if err != nil {
		http.Response(c, 400, "申请失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "申请成功", nil)
	return

}

func rechargeComplete(c *gin.Context) {
	var param model.RechargeParam

	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return

	}

	err := svr.RechargeSubmit(param)

	if err != nil {
		http.Response(c, 400, "申请失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "申请成功", nil)
	return

}
