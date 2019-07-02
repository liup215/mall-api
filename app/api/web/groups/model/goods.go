package model

import (
	xtime "mall/lib/time"
)

type EweiShopGroupsGoods struct {
	Id            int        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Displayorder  int        `json:"displayorder"`
	Uniacid       int        `json:"uniacid"`
	Title         string     `json:"title"`
	Category      int        `json:"category"`
	Stock         int        `json:"stock"`
	Price         float64    `json:"price"`
	Groupsprice   float64    `json:"groupsprice"`
	Singleprice   float64    `json:"singleprice"`
	Goodsnum      int        `json:"goodsnum"`
	Units         string     `json:"units"`
	Freight       float64    `json:"freight"`
	Endtime       xtime.Time `json:"endtime"`
	Groupnum      int        `json:"groupnum"`
	Sales         int        `json:"sales"`
	Thumb         string     `json:"thumb"`
	Description   string     `json:"description"`
	Content       string     `json:"content"`
	Createtime    xtime.Time `json:"createtime"`
	Status        int        `json:"status"`
	Ishot         int        `json:"ishot"`
	Deleted       int        `json:"deleted"`
	Goodsid       int        `json:"goodsid"`
	Followneed    int        `json:"followneed"`
	Followtext    string     `json:"followtext"`
	ShareTitle    string     `json:"share_title"`
	ShareIcon     string     `json:"share_icon"`
	ShareDesc     string     `json:"share_desc"`
	Goodssn       string     `json:"goodssn"`
	Productsn     string     `json:"productsn"`
	Showstock     int        `json:"showstock"`
	Purchaselimit int        `json:"purchaselimit"`
	Single        int        `json:"single"`
	Dispatchtype  int        `json:"dispatchtype"`
	Dispatchid    int        `json:"dispatchid"`
	Isindex       int        `json:"isindex"`
	Followurl     string     `json:"followurl"`
	Deduct        float64    `json:"deduct"`
	Rights        int        `json:"rights"`
	ThumbUrl      string     `json:"thumb_url"`
	Gid           int        `json:"gid"`
	Discount      int        `json:"discount"`
	Headstype     int        `json:"headstype"`
	Headsmoney    float64    `json:"headsmoney"`
	Headsdiscount int        `json:"headsdiscount"`
	Isdiscount    int        `json:"isdiscount"`
	Isverify      int        `json:"isverify"`
	Verifytype    int        `json:"verifytype"`
	Verifynum     int        `json:"verifynum"`
	Storeids      string     `json:"storeids"`
	Merchid       int        `json:"merchid"`
	Shorttitle    string     `json:"shorttitle"`
	Teamnum       int        `json:"teamnum"`
}

type GoodsQuery struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Uniacid  int    `json:"uniacid" form:"uniacid"`
	Type     string `json:"type" form:"type"`
	Keyword  string `json:"keyword" form:"keyword"`
}
