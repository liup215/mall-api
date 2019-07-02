package model

type PrePayParams struct {
	Uniacid int `json:"uniacid" form:"uniacid" binding:"required"`
	Orderid int `json:"orderid" form:"orderid" binding:"required"`
}
