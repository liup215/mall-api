package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/member/model"
	"mall/lib/time"
)

func (d *Dao) CreateMember(member model.EweiShopMember) error {
	member.Createtime = time.Now()
	return d.orm.Create(&member).Error
}

func (d *Dao) QueryMember(query model.MemberQuery) (model.EweiShopMember, error) {

	var u model.EweiShopMember
	err := d.parseQuery(query).First(&u).Error
	return u, err
}

func (d *Dao) parseQuery(query model.MemberQuery) *gorm.DB {

	var m model.EweiShopMember
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

	return d.orm.Where(&m)
}

func (d *Dao) UpdateMobile(id int, mobile string) error {
	return d.orm.Where(map[string]interface{}{"id": id}).Update("mobile", mobile).Error
}
