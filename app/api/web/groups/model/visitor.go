package model

import (
	xtime "mall/lib/time"
)

type EweiShopGroupsVisitor struct {
	Id          int        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid     int        `json:"id"`  // 应用ID
	Mid         int        `json:"mid"` // 用户ID
	Aid         string     `json:"aid"` // 活动ID
	Openid      string     `json:"openid"`
	Nickname    string     `json:"nickname"`
	Avatar      string     `json:"avatar"`
	Createdtime xtime.Time `json:"createdtime"`
}

type VisitorQuery struct {
	Id       int    `json:"id" form:"id"`
	Uniacid  int    `json:"uniacid" form:"uniacid"`
	Mid      int    `json:"mid" form:"mid"`
	Aid      string `json:"aid" form:"aid"`
	Openid   string `json:"openid" form:"openid"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
