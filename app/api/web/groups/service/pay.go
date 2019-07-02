package service

import (
	"errors"
	"mall/app/api/web/groups/model"

	// "mall/lib/strings"
	xtime "mall/lib/time"
)

func (s *Service) Prepay(openid string, p model.PrePayParams, clientIp string) (token string, err error) {
	if openid == "" {
		err = errors.New("无效的用户Openid")
		return
	}
	if p.Uniacid == 0 {
		err = errors.New("无效的应用ID")
		return
	}

	if p.Orderid == 0 {
		err = errors.New("无效的订单ID")
		return
	}

	var order *model.EweiShopGroupsOrder
	if order, err = s.d.QueryOrderOne(model.OrderQuery{Id: p.Orderid}); err != nil {
		return
	}

	if order.Openid != openid {
		err = errors.New(("没有订单权限"))
		return
	}

	if order.IsTeam == 1 && order.Success == 1 {
		err = errors.New("订单已失效，请浏览其他商品或联系商家")
		return
	}

	// 活动有效期验证
	if order.Aid != "" {
		var act *model.EweiShopGroupsActivity
		if act, err = s.d.QueryActivityById(order.Aid); err != nil {
			return
		}

		if act.Endtime < xtime.Now() {
			err = errors.New("活动已过期，下次记得早点儿来哦！")
			return
		}

		var teamOrders []model.EweiShopGroupsOrder
		if teamOrders, _, err = s.d.QueryOrderList(model.OrderQuery{Uniacid: p.Uniacid, Teamid: order.Teamid}); err != nil {
			return
		}

		for _, o := range teamOrders {
			if o.Success == -1 {
				err = errors.New("该活动已过期，请参与其他活动或联系商家")
				return
			}

			if o.Success == 1 {
				err = errors.New("该活动已结束，请参与其他活动或联系商家")
			}
		}

		var num int
		num, err = s.d.OrderCount(model.OrderQuery{Uniacid: p.Uniacid, Teamid: order.Teamid, Statuses: []int{1, 2, 3}})
		if order.Groupnum <= num {
			err = errors.New("该活动已成功组团，请浏览其他活动")
		}
	}

	if order.Status == -1 || order.Status >= 1 {
		err = errors.New("订单已关闭")
		return
	}

	if order.Status >= 1 {
		err = errors.New("订单已支付")
		return
	}

	return s.d.PrePayToken(*order, clientIp)

}
