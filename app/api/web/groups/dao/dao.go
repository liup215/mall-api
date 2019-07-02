package dao

import (
	"mall/app/api/web/groups/conf"
	"mall/app/api/web/groups/model"
	"mall/lib/database/nosql/mongo"
	"mall/lib/database/orm"
	"mall/lib/database/redis"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	orm    *gorm.DB
	mgo    *mongo.Mongo
	redigo *redis.Redis
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	if !db.HasTable(&model.EweiShopGroupsVisitor{}) {
		db.CreateTable(&model.EweiShopGroupsVisitor{})
	}
	return &Dao{
		orm:    db,
		mgo:    mongo.New(c.Mongo),
		redigo: redis.New(c.Redis),
	}
}
