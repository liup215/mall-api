package server

import (
	"github.com/gin-gonic/gin"
	"mall/app/service/main/finance/model"
	"mall/lib/net/http"
)

func preRecharge(c *gin.Context) {
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

func rechargeLogConfirm(c *gin.Context) {
	var param model.RechargeLogConfirmParam

	if err := c.Bind(&param); err != nil {
		http.Response(c, 400, "参数错误,"+err.Error(), nil)
		return
	}

	res, err := svr.RechargeLogConfirm(param)

	if err != nil {
		http.Response(c, 400, "确认失败,"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "确认成功", res)
	return

}
