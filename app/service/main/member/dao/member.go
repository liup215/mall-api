package dao

import (
	"mall/app/service/main/member/model"
)

func (d *Dao) QueryMember(query model.MemberQuery) (model.EweiShopMember, error) {
	var m, u model.EweiShopMember
	if query.Id != 0 {
		m.Id = query.Id
	}

	if query.Uniacid != 0 {
		m.Uniacid = query.Uniacid
	}

	if query.Openid != "" {
		m.Openid = query.Openid
	}

	if query.Mobile != "" {
		m.Mobile = query.Mobile
	}

	err := d.db.Where(&m).First(&u).Error
	return u, err
}
