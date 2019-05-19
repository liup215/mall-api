package model

type Page struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page" form:"pageSize"`
}

type LogQuery struct {
	Id      int    `json:"id" form:"id" uri:"id"`
	Uniacid int    `json:"uniacid" form:"uniacid" uri:"uniacid"`
	Openid  string `json:"openid" form:"openid" uri:"openid"`
	Type    int    `json:"-" form:"-" url:"-"`
	TypeStr string `json:"type" form:"type" uri:"type"`
}
