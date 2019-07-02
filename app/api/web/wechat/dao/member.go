package dao

import (
	"errors"
	"mall/app/api/web/wechat/model"
	xtime "mall/lib/time"

	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateMcMembers(m model.McMembers) (*model.McMembers, error) {
	if m.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if m.Uid != 0 {
		m.Uid = 0
	}

	m.Createtime = xtime.Now()

	err := d.orm.Create(&m).Error
	return &m, err
}

func (d *Dao) QueryMcMemberOne(q model.McMembersQuery) (*model.McMembers, error) {
	db := d.parseMcMembersQuery(q)

	var m model.McMembers
	err := db.First(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, err
}

func (d *Dao) parseMcMembersQuery(q model.McMembersQuery) *gorm.DB {
	db := d.orm.Model(&model.McMembers{})

	if q.Uniacid != 0 {
		db = db.Where("uniacid = ?", q.Uniacid)
	}

	if q.Uid != 0 {
		db = db.Where("uid = ?", q.Uid)
	}

	return db
}
