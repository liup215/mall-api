package model

import (
	xtime "mall/lib/time"
)

type McMembers struct {
	Uid             int        `gorm:"primary_key;AUTO_INCREMENT" json:"uid"`
	Uniacid         int        `json:"uniacid"`
	Mobile          string     `json:"mobile"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	Salt            string     `json:"salt"`
	Groupid         int        `json:"groupid"`
	Credit1         float64    `json:"credit1"`
	Credit2         float64    `json:"credit2"`
	Credit3         float64    `json:"credit3"`
	Credit4         float64    `json:"credit4"`
	Credit5         float64    `json:"credit5"`
	Credit6         float64    `json:"credit6"`
	Createtime      xtime.Time `json:"createtime"`
	Realname        string     `json:"realname"`
	Nickname        string     `json:"nickname"`
	Avatar          string     `json:"avatar"`
	Qq              string     `json:"qq"`
	Vip             int        `json:"vip"`
	Gender          int        `json:"gender"`
	Birthyear       int        `json:"birthyear"`
	Birthmonth      int        `json:"birthmonth"`
	Birthday        int        `json:"birthday"`
	Constellation   string     `json:"constellation"`
	Zodiac          string     `json:"zodiac"`
	Telephone       string     `json:"telephone"`
	Idcard          string     `json:"idcard"`
	Studentid       string     `json:"studentid"`
	Grade           string     `json:"grade"`
	Address         string     `json:"address"`
	Zipcode         string     `json:"zipcode"`
	Nationality     string     `json:"nationality"`
	Resideprovince  string     `json:"resideprovince"`
	Residecity      string     `json:"residecity"`
	Residedist      string     `json:"residedist"`
	Graduateschool  string     `json:"graduateschool"`
	Company         string     `json:"company"`
	Education       string     `json:"education"`
	Occupation      string     `json:"occupation"`
	Position        string     `json:"position"`
	Revenue         string     `json:"revenue"`
	Affectivestatus string     `json:"affectivestatus"`
	Lookingfor      string     `json:"lookingfor"`
	Bloodtype       string     `json:"bloodtype"`
	Height          string     `json:"height"`
	Weight          string     `json:"weight"`
	Alipay          string     `json:"alipay"`
	Msn             string     `json:"msn"`
	Taobao          string     `json:"taobao"`
	Site            string     `json:"site"`
	Bio             string     `json:"bio"`
	Interest        string     `json:"interest"`
}

type McMembersQuery struct {
	Uniacid int `json:"uniacid" form:"uniacid"`
	Uid     int `json:"uid" form:"uid"`
}
