package model

import (
	"mall/app/service/main/adv/api"
)

type ShopAdv struct {
	Id           int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid      int    `json:"uniacid"`
	Advname      string `json:"advname"`
	Link         string `json:"link"`
	Thumb        string `json:"thumb"`
	Displayorder int    `json:"displayorder"`
	Enabled      int    `json:"enabled"`
	Shopid       int    `json:"shopid"`
	Iswxapp      int    `json:"iswxapp"`
}

func (*ShopAdv) TableName() string {
	return "ewei_shop_adv"
}

func (adv *ShopAdv) ToPbAdv() *api.Adv {
	return &api.Adv{
		Id:           int64(adv.Id),
		Uniacid:      int64(adv.Uniacid),
		Advname:      adv.Advname,
		Link:         adv.Link,
		Thumb:        adv.Link,
		Displayorder: int64(adv.Displayorder),
		Enabled:      int64(adv.Enabled),
		Shopid:       int64(adv.Shopid),
		Iswxapp:      int64(adv.Iswxapp),
	}
}
