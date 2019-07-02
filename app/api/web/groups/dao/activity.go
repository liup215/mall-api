package dao

import (
	"errors"
	"mall/app/api/web/groups/model"
	xtime "mall/lib/time"

	"gopkg.in/mgo.v2/bson"
)

func (d *Dao) InsertActivity(act model.EweiShopGroupsActivity) (*model.EweiShopGroupsActivity, error) {
	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()
	act.ID = bson.NewObjectId()
	act.Id = act.ID.Hex()
	act.Createdtime = xtime.Now()
	err := adb.C.Insert(&act)
	if err != nil {
		return nil, err
	}

	return &act, nil
}

func (d *Dao) UpdateActivity(act model.EweiShopGroupsActivity) (*model.EweiShopGroupsActivity, error) {
	if act.Id == "" {
		return nil, errors.New("aid不能为空")
	}

	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()
	err := adb.C.Update(bson.M{"id": act.Id}, bson.M{"$set": &act})
	return &act, err
}

func (d *Dao) IncrActivityVisits(id string) error {
	if id == "" {
		return errors.New("aid不能为空")
	}
	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()

	return adb.C.Update(bson.M{"id": id}, bson.M{"$inc": bson.M{"visit": 1}})
}

func (d *Dao) IncrActivityShares(id string) error {
	if id == "" {
		return errors.New("aid不能为空")
	}
	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()

	return adb.C.Update(bson.M{"id": id}, bson.M{"$inc": bson.M{"shares": 1}})
}

func (d *Dao) QueryActivityById(aid string) (*model.EweiShopGroupsActivity, error) {
	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()
	var act model.EweiShopGroupsActivity
	err := adb.C.Find(bson.M{"id": aid}).One(&act)
	return &act, err
}

func (d *Dao) QueryActivityList(q model.ActivityQuery) ([]model.EweiShopGroupsActivity, int) {
	adb := d.mgo.Model("ewei_shop_groups_activity")
	defer adb.Close()

	if q.Page <= 0 {
		q.Page = 1
	}

	if q.PageSize <= 0 {
		q.PageSize = 20
	}

	m := d.parseActivityQuery(q)

	finds := adb.C.Find(m)

	var list []model.EweiShopGroupsActivity
	finds.Sort("-createdtime").Skip((q.Page - 1) * q.PageSize).Limit(q.PageSize).All(&list)

	total, _ := finds.Count()

	return list, total
}

// func (d *Dao) QueryActivityOne(q model.ActivityQuery) (*model.EweiShopGroupsActivity, error) {
// 	adb := d.mgo.Model("ewei_shop_groups_activity")
// 	defer adb.Close()

// 	m := d.parseActivityQuery(q)

// 	fmt.Println(m)
// 	var act model.EweiShopGroupsActivity
// 	err := adb.C.Find(m).One(&act)

// 	return &act, err
// }

func (d *Dao) parseActivityQuery(q model.ActivityQuery) bson.M {

	c := bson.M{}
	if q.Id != "" {
		c["id"] = q.Id
	}

	if q.Uniacid != 0 {
		c["uniacid"] = q.Uniacid
	}

	if q.Uid != 0 {
		q.Uids = append(q.Uids, q.Uid)
	}

	if len(q.Uids) > 0 {
		c["uids"] = bson.M{
			"$in": q.Uids,
		}
	}

	if q.Keyword != "" {
		c["title"] = bson.M{
			"$regex": bson.RegEx{q.Keyword, "i"},
		}
	}

	if q.Status != -1 {
		q.Statuses = append(q.Statuses, q.Status)
	}

	if len(q.Statuses) > 0 {
		c["status"] = bson.M{
			"$in": q.Statuses,
		}
	}

	return c
}
