package server

import (
	"mall/app/service/main/payment/server/rpc/client"
	"mall/lib/net/http"

	"github.com/gin-gonic/gin"
)

func pay(c *gin.Context) {
	var params client.UnifiedOrderParams
	if err := c.Bind(&params); err != nil {
		http.Response(c, 400, "参数错误"+err.Error(), nil)
		return
	}

	params.NotifyUrl = c.Request.Host + "/notify"

	reply, err := svr.Pay(params)
	if err != nil {
		http.Response(c, 400, "支付失败"+err.Error(), nil)
		return
	}

	http.Response(c, 200, "支付成功", reply)

}
