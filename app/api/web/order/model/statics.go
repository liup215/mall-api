package model

type OrderStaticsParam struct {
	Uniacid int    `json:"uniacid" form:"uniacid" binding:"required"`
	Openid  string `json:"openid" form:"openid"`
	Status  int    `json:"status" form:"status"`
}

type OrderStaticsResponse struct {
	Count int     `json:"count"`
	Money float64 `json:"money"`
}
