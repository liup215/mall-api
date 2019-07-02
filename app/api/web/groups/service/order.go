package service

import (
	"errors"
	"mall/app/api/web/groups/model"
	xtime "mall/lib/time"

	"github.com/jinzhu/gorm"
)

func (s *Service) ConfirmOrder(order model.EweiShopGroupsOrder) (*model.EweiShopGroupsOrder, error) {
	if order.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}
	if order.Openid == "" {
		return nil, errors.New("无效的用户openid")
	}

	// 活动相关判断，如果aid不是空，说明当前是活动拼团
	if order.Aid != "" {
		if err := s.handleActivityOrder(&order); err != nil {
			return nil, err
		}
	}

	return s.d.CreateOrder(order)
}

func (s *Service) handleActivityOrder(order *model.EweiShopGroupsOrder) error {
	act, err := s.d.QueryActivityById(order.Aid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("查无此活动")
		} else {
			return errors.New("发生未知错误 ," + err.Error())
		}
	}

	now := xtime.Now()
	if now < act.Starttime {
		return errors.New("活动尚未开始")
	}

	if now > act.Endtime {
		return errors.New("活动已经结束")
	}

	ordernum, err := s.d.OrderCount(model.OrderQuery{Aid: order.Aid, Openid: order.Openid, Statuses: []int{0, 1, 2, 3}, Uniacid: order.Uniacid})
	if act.Purchaselimit != 0 && act.Purchaselimit <= ordernum {
		return errors.New("您已到达此活动上限，请参与其他活动或联系主办方！")
	}
	_, err = s.d.QueryOrderOne(model.OrderQuery{Aid: order.Aid, Openid: order.Openid, Status: 0, Uniacid: order.Uniacid})
	if err != gorm.ErrRecordNotFound {
		if err == nil {
			return errors.New("您有未支付订单，请尽快支付")
		} else {
			return errors.New("发生未知错误, " + err.Error())
		}
	}

	// 检查团购相关
	if order.IsTeam == 1 {
		if order.Teamid != 0 {
			// 参加团
			orders, total, err := s.d.QueryOrderList(model.OrderQuery{Aid: order.Aid, Statuses: []int{1, 2, 3}, Teamid: order.Teamid, Uniacid: order.Uniacid, Page: 1, PageSize: order.Groupnum})
			if err != nil {
				return errors.New("获取团购订单失败, " + err.Error())
			}

			if total == 0 {
				return errors.New("无效的团")
			}

			if orders[0].Groupnum <= total {
				return errors.New("该团已成功组团，请浏览其他商品或联系商家")
			}

			for _, o := range orders {
				if o.Success == 1 {
					return errors.New("活动已结束，请浏览其他商品或联系商家")
				}
				if o.Success == -1 {
					return errors.New("活动已过期，请浏览其他商品或联系商家")
				}
			}

			order.Price = orders[0].Price
			order.Groupnum = orders[0].Groupnum
		} else {
			// 新建团，需要判断团购价格和团购人数是否是活动中指定的
			flag := false
			for _, teamset := range act.Teamset {
				if teamset.TeamNum == order.Groupnum && teamset.TeamPrice == order.Price {
					flag = true
					break
				}
			}
			// 如果flag为false，说明活动中没有这个设置
			if !flag {
				return errors.New("错误的开团人数和团购价格设置")
			}
		}
	} else {
		order.Price = act.Singleprice
	}

	// 地址设置
	if order.Mobile == "" || order.Realname == "" {
		return errors.New("姓名和联系方式不能为空")
	}

	return nil
}

func (s *Service) UnpayedOrder(uniacid int, openid, aid string) (*model.EweiShopGroupsOrder, error) {
	return s.d.QueryOrderOne(model.OrderQuery{Openid: openid, Status: 0, Aid: aid})
}

func (s *Service) QueryActBuyerList(actId string) ([]model.EweiShopGroupsBuyer, int) {
	var list []model.EweiShopGroupsBuyer
	total, _ := s.d.OrderCount(model.OrderQuery{Aid: actId, Status: 1})
	orders, _, _ := s.d.QueryOrderList(model.OrderQuery{Aid: actId, Status: 1, PageSize: total})

	for _, o := range orders {
		buyer := model.EweiShopGroupsBuyer{
			Openid:  o.Openid,
			Payedat: o.Paytime,
			Price:   o.Price,
			Name:    o.Realname,
			Avatar:  o.Avatar,
		}
		list = append(list, buyer)
	}
	return list, total
}
