package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/payment"
)

func notify(c *gin.Context) {
	err := payment.HandlePaidNotify(c.Writer, c.Request, func(ntf payment.PaidNotify) (bool, string) {
		// 先判断订单状态是否更新 如果 == 1 不更新
		// 如果 != 1 更新
		plidStr := ntf.Attach
		plid, err := strconv.Atoi(plidStr)
		if err != nil {
			return false, "plogID无效"
		}

		err = svr.PayComplete(plid)
		if err != nil {
			return false, ""
		}

		return true, ""
	})
	if err != nil {
		// 这是支付不成功的
		fmt.Println(err)
		return
	}
}
