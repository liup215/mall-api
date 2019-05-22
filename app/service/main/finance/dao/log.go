package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/finance/model"
	"mall/lib/strings"
	"mall/lib/time"
)

func (d *Dao) CreateLog(log model.EweiShopMemberLog) (model.EweiShopMemberLog, error) {

	logno := ""
	for {
		logno = strings.CreateNo("RC")

		var record model.EweiShopMemberLog
		if d.orm.Where(&model.EweiShopMemberLog{Logno: logno}).First(&record).Error == gorm.ErrRecordNotFound {
			break
		}
	}

	log.Logno = logno
	log.Createtime = time.Now()

	err := d.orm.Create(&log).Error
	return log, err

}

func (d *Dao) LogByLogno(logno string, uniacid int) (model.EweiShopMemberLog, error) {
	var log model.EweiShopMemberLog

	err := d.orm.Where(map[string]interface{}{"logno": logno, "uniacid": uniacid}).First(&log).Error
	return log, err
}

func (d *Dao) LogById(id, uniacid int) (model.EweiShopMemberLog, error) {
	var log model.EweiShopMemberLog

	err := d.orm.Where(map[string]interface{}{"id": id, "uniacid": uniacid}).First(&log).Error
	return log, err
}

func (d *Dao) LogList(q model.LogQuery, page model.Page) ([]model.EweiShopMemberLog, int, error) {
	var list []model.EweiShopMemberLog
	var total int
	db := d.parseLogQuery(q)

	if page.Page <= 0 {
		page.Page = 1
	}

	if page.PageSize <= 0 {
		page.PageSize = 10
	}
	err := db.Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}

	db.Count(&total)

	return list, total, err
}

func (d *Dao) QueryLog(q model.LogQuery) (model.EweiShopMemberLog, error) {
	var log model.EweiShopMemberLog
	db := d.parseLogQuery(q)
	err := db.First(&log).Error

	return log, err
}

func (d *Dao) parseLogQuery(q model.LogQuery) *gorm.DB {

	var l model.EweiShopMemberLog
	if q.Id != 0 {
		l.Id = q.Id
	}

	if q.Uniacid != 0 {
		l.Uniacid = q.Uniacid
	}

	if q.Openid != "" {
		l.Openid = q.Openid
	}

	if q.Logno != "" {
		l.Logno = q.Logno
	}

	db := d.orm.Model(&model.EweiShopMemberLog{}).Where(&l)

	switch q.TypeStr {
	case "all":
		q.Type = -1
	case "credit1":
		q.Type = 1
	case "credit2":
		q.Type = 0
	default:
		q.Type = -1
	}

	if q.Type != -1 {
		db = db.Where("type = ?", q.Type)
	}
	return db

}

func (d *Dao) UpdateLogStatus(logno string, status, uniacid int) error {
	return d.orm.Model(&model.EweiShopMemberLog{}).Where(map[string]interface{}{"logno": logno, "uniacid": uniacid}).Update("status", status).Error
}
