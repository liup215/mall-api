package model

import (
	"mall/lib/address"
	xtime "mall/lib/time"

	"gopkg.in/mgo.v2/bson"
)

type EweiShopGroupsActivity struct {
	ID             bson.ObjectId                   `bson:"_id,omitempty" gorm:"-" json:"-"`
	Id             string                          `bson:"id,omitempty" json:"id"`
	Uniacid        int                             `bson:"uniacid,omitempty" json:"uniacid"`
	Uid            int                             `bson:"uid,omitempty" json:"uid"`
	Title          string                          `bson:"title,omitempty" json:"title"`
	Isbullet       int                             `bson:"isbullet,omitempty" json:"isbullet"` // 是否显示弹幕
	Starttime      xtime.Time                      `bson:"starttime,omitempty" json:"starttime"`
	Endtime        xtime.Time                      `bson:"endtime,omitempty" json:"endtime"`
	ThumbUrl       string                          `bson:"thumb_url,omitempty" json:"thumb_url"`
	ShareTitle     string                          `bson:"share_title,omitempty" json:"share_title"`
	ShareImg       string                          `bson:"share_img,omitempty" json:"share_img"`
	ShareDesc      string                          `bson:"share_desc,omitempty" gorm:"type:TEXT" json:"share_desc"`
	Content        string                          `bson:"content,omitempty" gorm:"type:TEXT" json:"content"`
	Visit          int                             `bson:"visit,omitempty" json:"visit"` // 访问次数
	Sales          int                             `bson:"sales,omitempty" json:"sales"` // 购买数量
	Shares         int                             `bson:"shares,omitempty" json:"shares"`
	Purchaselimit  int                             `bson:"purchase_limit,omitempty" json:"purchaselimit"`
	Price          float64                         `bson:"price,omitempty" json:"price"`                // 团购原价
	Singleprice    float64                         `bson:"singleprice,omitempty" json:"singleprice"`    // 单买价格
	Groupnum       int                             `bson:"groupnum,omitempty" json:"groupnum"`          // 开团人数
	ServerPhone    string                          `bson:"server_phone,omitempty" json:"service_phone"` // 客服电话
	ServerName     string                          `bson:"server_name,omitempty" json:"server_name"`    // 客服姓名
	Status         int                             `bson:"status,omitempty" json:"status"`              // 状态
	Createdtime    xtime.Time                      `bson:"createdtime,omitempty" json:"createtime"`     // 创建时间
	GoodsListArray []int                           `bson:"goods_list_array,omitempty" gorm:"-" json:"goods_list"`
	VerifyCode     string                          `bson:"verify_code,omitempty" json:"verify_code"` // 核销密码
	Isteam         int                             `bson:"isteam,omitempty" json:"isteam"`
	Teamset        []EweiShopGroupsActivityTeamset `bson:"teamset,omitempty" json:"teamset" grom:"ForeignKey:ActivityId"`
	Address        address.Address                 `bson:"address,omitempty" json:"address,omitempty"`
}

type EweiShopGroupsActivityTeamset struct {
	Name      string  `bson:"name" json:"name"`
	TeamNum   int     `bson:"team_num" json:"teamnum"`      // 成团人数
	TeamPrice float64 `bson:"team_price" json:"team_price"` // 团购价格
}

type ActivityQuery struct {
	Id       string `json:"id" form:"id"`
	Uniacid  int    `json:"uniacid" form:"uniacid"`
	Uid      int    `json:"uid" form:"uid"`
	Uids     []int  `json:"uids" form:"uids"`
	Status   int    `json:"status" form:"status"`
	Statuses []int  `json:"statuses" form:"statuses"`
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
