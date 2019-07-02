package model

type McGroups struct {
	Groupid   int    `gorm:"primary_key;AUTO_INCREMENT" json:"groupid"`
	Uniacid   int    `json:"uniacid"`
	Title     string `json:"title"`
	Credit    int    `json:"credit"`
	Isdefault int    `json:"isdefault"`
	Orderlist int    `json:"orderlist"`
}

type McGroupsQuery struct {
	Uniacid   int `json:"uniacid"`
	Isdefault int `json:"isdefault"`
}
