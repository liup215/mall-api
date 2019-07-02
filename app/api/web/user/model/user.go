package model

import (
	"mall/lib/time"
)

type Users struct {
	Uid            int       `gorm:"primary_key;AUTO_INCREMENT" json:"uid"`
	OwnerUid       int       `json:"owner_uid"`
	Groupid        int       `json:"groupid"`
	FounderGroupid int       `json:"founder_groupid"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Salt           string    `json:"-"`
	Type           int       `json:"type"`
	Status         int       `json:"status"`
	Joindate       time.Time `json:"joindate"`
	Joinip         string    `json:"joinip"`
	Lastvisit      int       `json:"lastvisit"`
	Lastip         string    `json:"lastip"`
	Remark         string    `json:"remark"`
	Starttime      time.Time `json:"starttime"`
	Endtime        time.Time `json:"endtime"`
	Uniacid        int       `json:"uniacid"`
	Tid            int       `json:"tid"`
	Schoolid       int       `json:"schoolid"`
	RegisterType   int       `json:"register_type"`
	Openid         string    `json:"openid"`
}

type UserQuery struct {
	Uid      int
	Username string
	Openid   string
}
