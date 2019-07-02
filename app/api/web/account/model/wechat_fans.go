package model

import (
	xtime "mall/lib/time"
)

type Mc_mapping_fans struct {
	Fanid        int        `gorm:"primary_key;AUTO_INCREMENT" json:"fanid"`
	Acid         int        `json:"acid"`
	Uniacid      int        `json:"uniacid"`
	Uid          int        `json:"uid"`
	Openid       string     `json:"openid"`
	Nickname     string     `json:"nickname"`
	Groupid      string     `json:"groupid"`
	Salt         string     `json:"salt"`
	Follow       int        `json:"follow"`
	Followtime   xtime.Time `json:"followtime"`
	Unfollowtime xtime.Time `json:"unfollowtime"`
	Tag          string     `json:"tag"`
	Updatetime   xtime.Time `json:"updatetime"`
	Unionid      string     `json:"unionid"`
}
