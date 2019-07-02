package model

import (
	xtime "mall/lib/time"
)

type McMappingFans struct {
	Fanid        int        `gorm:"primary_key;AUTO_INCREMENT" json:"fanid"`
	Acid         int        `json:"acid"`
	Uniacid      int        `json:"uniacid"`
	Uid          int        `json:"uid"`
	Openid       string     `json:"openid"`
	Nickname     string     `json:"nickname"`
	Groupid      string     `json:"groupid"`
	Avatar       string     `json:"avatar"`
	Salt         string     `json:"-"`
	Follow       int        `json:"follow"`
	Followtime   xtime.Time `json:"followtime"`
	Unfollowtime xtime.Time `json:"unfollowtime"`
	Tag          string     `json:"tag"`
	Updatetime   xtime.Time `json:"updatetime"`
	Unionid      string     `json:"unionid"`
}

type UserInfoParams struct {
	Uniacid int    `json:"uniacid" form:"uniacid"`
	Code    string `json:"code" form:"code"`
	State   string `json:"state" form:"state"`
}

type FansQuery struct {
	Uniacid int    `json:"uniacid" form:"uniacid"`
	Fanid   int    `json:"fanid" form:"fanid"`
	Openid  string `json:"openid" form:"openid"`
}
