package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/member/model"
	"mall/lib/time"
)

func (d *Dao) CreateMember(member model.EweiShopMember) (*model.EweiShopMember, error) {
	member.Createtime = time.Now()
	err := d.orm.Create(&member).Error
	return &member, err
}

func (d *Dao) QueryMember(query model.MemberQuery) (model.EweiShopMember, error) {

	var u model.EweiShopMember
	err := d.parseQuery(query).First(&u).Error
	return u, err
}

func (d *Dao) QueryMemberOne(query model.MemberQuery) (*model.EweiShopMember, error) {

	var u model.EweiShopMember
	err := d.parseQuery(query).First(&u).Error
	return &u, err
}

func (d *Dao) parseQuery(query model.MemberQuery) *gorm.DB {
	db := d.orm.Model(&model.EweiShopMember{})

	if query.Id != 0 {
		db.Where("id = ?", query.Id)
	}

	if query.Uniacid != 0 {
		db.Where("uniacid = ?", query.Uniacid)
	}

	if query.Openid != "" {
		db.Where("openid = ?", query.Openid)
	}

	if query.Mobile != "" {
		db.Where("mobile = ? AND mobileverify = ?", query.Mobile, 1)
	}

	return db
}

func (d *Dao) UpdateMobile(id int, mobile string) error {
	return d.orm.Model(&model.EweiShopMember{}).Where(map[string]interface{}{"id": id}).Update("mobile", mobile).Error
}
