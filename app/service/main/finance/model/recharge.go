package model

type RechargeParam struct {
	Money   float64 `json:"money" form:"money"`
	Uniacid int     `json:"uniacid" form:"uniacid"`
	Openid  string  `json:"openid" form:"openid"`
}
