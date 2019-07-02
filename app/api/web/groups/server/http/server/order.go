package server

import (
	"mall/app/api/web/groups/model"
	"mall/lib/net/http"
	"mall/lib/net/http/middleware/auth"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func confirmOrder(c *gin.Context) {
	u, exist := auth.GetCurrentUser(c)
	if !exist {
		http.Response(c, 400, "用户解析失败", nil)
		return
	}

	if u.Usertype == 0 {
		http.Response(c, 400, "无效的用户", nil)
		return
	}

	var order model.EweiShopGroupsOrder
	if err := c.Bind(&order); err != nil {
		http.Response(c, 400, "参数错误"+err.Error(), nil)
		return
	}

	order.Openid = u.Openid
	order.Avatar = u.Avatar
	o, err := svr.ConfirmOrder(order)
	if err != nil {
		http.Response(c, 400, err.Error(), nil)
		return
	}

	http.Response(c, 200, "订单创建成功", o)
}

func unpayedOrder(c *gin.Context) {
	u, exist := auth.GetCurrentUser(c)
	if !exist {
		http.Response(c, 400, "用户解析失败", nil)
		return
	}

	if u.Usertype == 0 {
		http.Response(c, 400, "无效的用户", nil)
		return
	}
	aid := c.Query("aid")

	o, err := svr.UnpayedOrder(u.Uniacid, u.Openid, aid)

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			http.Response(c, 400, err.Error(), nil)
			return
		} else {
			http.Response(c, 200, "没有未支付订单", nil)
			return
		}
	}

	http.Response(c, 200, "订单获取成功", o)
}
