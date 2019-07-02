package model

import (
	"mall/lib/time"
)

type EweiShopMemberLog struct {
	Id             int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid        int       `json:"uniacid,omitempty"`
	Openid         string    `json:"openid,omitempty"`
	Type           int       `json:"type,omitempty"`
	Logno          string    `json:"logno,omitempty"`
	Title          string    `json:"title,omitempty"`
	Createtime     time.Time `json:"createtime,omitempty"`
	Status         int       `json:"status,omitempty"`
	Money          float64   `json:"money,omitempty"`
	Rechargetype   string    `json:"rechargetype,omitempty"`
	Gives          float64   `json:"gives,omitempty"`
	Couponid       int       `json:"couponid,omitempty"`
	Transid        string    `json:"transid,omitempty"`
	Realmoney      float64   `json:"realmoney,omitempty"`
	Charge         float64   `json:"charge,omitempty"`
	Deductionmoney float64   `json:"deductionmoney,omitempty"`
	Isborrow       int       `json:"isborrow,omitempty"`
	Borrowopenid   string    `json:"borrowopenid,omitempty"`
	Remark         string    `json:"remark,omitempty"`
	Apppay         int       `json:"apppay,omitempty"`
	Alipay         string    `json:"alipay,omitempty"`
	Bankname       string    `json:"bankname,omitempty"`
	Bankcard       string    `json:"bankcard,omitempty"`
	Realname       string    `json:"realname,omitempty"`
	Applytype      int       `json:"applytype,omitempty"`
	Sendmoney      float64   `json:"sendmoney,omitempty"`
	Senddata       string    `json:"senddata,omitempty"`
}

type EweiShopMemberLogQuery struct {
	Id      int    `json:"id" form:"id"`
	Uniacid int    `json:"uniacid" form:"uniacid"`
	Openid  string `json:"openid" form:"openid"`
	Type    int    `json:"type" form:"type"`
	Status  int    `json:"status" form:"status"`
	Logno   string `json:"logno" form:"logno"`
}
