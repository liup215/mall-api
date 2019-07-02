package dao

import (
	"errors"
	"mall/app/api/web/groups/model"
	paymentModel "mall/app/service/main/payment/model"
	"mall/lib/strings"
)

func (d *Dao) PrePayToken(order model.EweiShopGroupsOrder, clientIp string) (string, error) {
	token := strings.Md5("prepay:" + order.Orderno)

	if order.Uniacid == 0 {
		return "", errors.New("订单uniacid为空")
	}
	o := paymentModel.PrepayOrder{
		Orderno:  order.Orderno,
		Module:   "groups",
		Uniacid:  order.Uniacid,
		Openid:   order.Openid,
		Fee:      order.Price,
		Body:     "E贸平台-活动参与",
		CreateIp: clientIp,
	}
	err := d.redigo.HMSET(token, &o, 2*60)
	if err != nil {
		return "", err
	}
	return token, nil
}
