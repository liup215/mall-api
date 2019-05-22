package model

type PreRechargeParam struct {
	Money   float64 `json:"money" form:"money"`
	Uniacid int     `json:"uniacid" form:"uniacid"`
	Openid  string  `json:"openid" form:"openid"`
	Logno   string  `json:"logno" form:"logno"`
}

type PreRechargeResponse struct {
	Uniacid int    `json:"uniacid"`
	Openid  string `json:"openid"`
	Logno   string `json:"logno"`
}

type RechargeLogConfirmParam struct {
	Money   float64 `json:"money" form:"money"`
	Uniacid int     `json:"uniacid" form:"uniacid"`
	Openid  string  `json:"openid" form:"openid"`
	Logno   string  `json:"logno" form:"logno"`
}

type RechargeLogConfirmResponse struct {
	Uniacid int     `json:"uniacid"`
	Logno   string  `json:"logno"`
	Openid  string  `json:"openid"`
	Money   float64 `json:"money"`
	Type    int     `json:"type"`
}
