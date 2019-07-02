package model

import (
	xtime "mall/lib/time"
)

type EweiShopGroupsOrder struct {
	Id           int        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid      int        `json:"uniacid"`
	Openid       string     `json:"openid"`
	Avatar       string     `json:"avatar"`
	Aid          string     `json:"aid"`
	Orderno      string     `json:"orderno"`
	Groupnum     int        `json:"groupnum"`
	Paytime      xtime.Time `json:"paytime"`
	Price        float64    `json:"price"`
	Freight      float64    `json:"freight"`
	Status       int        `json:"status"`
	PayType      string     `json:"pay_type"`
	Goodid       int        `json:"goodid"`
	Teamid       int        `json:"teamid"`
	IsTeam       int        `json:"is_team"`
	Heads        int        `json:"heads"`
	Starttime    xtime.Time `json:"starttime"`
	Endtime      xtime.Time `json:"endtime"`
	Createtime   xtime.Time `json:"createtime"`
	Success      int        `json:"success"`
	Delete       int        `json:"delete"`
	Credit       int        `json:"credit"`
	Creditmoney  float64    `json:"creditmoney"`
	Dispatchid   int        `json:"dispatchid"`
	Addressid    int        `json:"addressid"`
	Address      string     `json:"address"`
	Discount     float64    `json:"discount"`
	Canceltime   xtime.Time `json:"canceltime"`
	Finishtime   xtime.Time `json:"finishtime"`
	Refundid     int        `json:"refundid"`
	Refundstate  int        `json:"refundstate"`
	Refundtime   xtime.Time `json:"refundtime"`
	Express      string     `json:"express"`
	Expresscom   string     `json:"expresscom"`
	Expresssn    string     `json:"expresssn"`
	Sendtime     xtime.Time `json:"sendtime"`
	Remark       string     `json:"remark"`
	Remarkclose  string     `json:"remarkclose"`
	Remarksend   string     `json:"remarksend"`
	Message      string     `json:"message"`
	Deleted      int        `json:"deleted"`
	Realname     string     `json:"realname"`
	Mobile       string     `json:"mobile"`
	Isverify     int        `json:"isverify"`
	Verifytype   int        `json:"verifytype"`
	Verifycode   string     `json:"verifycode"`
	Verifynum    int        `json:"verifynum"`
	Printstate   int        `json:"printstate"`
	Printstate2  int        `json:"printstate2"`
	Apppay       int        `json:"apppay"`
	Isborrow     int        `json:"isborrow"`
	Borrowopenid string     `json:"borrowopenid"`
}

type OrderQuery struct {
	Id       int    `json:"id" form:"id"`
	Uniacid  int    `json:"uniacid" form:"uniacid"`
	Openid   string `json:"openid" form:"openid"`
	Aid      string `json:"aid" form:"aid"`
	Orderno  string `json:"orderno" form:"orderno"`
	Status   int    `json:"status" form:"status"`
	Statuses []int  `json:"statuses" form:"statuses"`
	Teamid   int    `json:"teamid" form:"teamid"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
