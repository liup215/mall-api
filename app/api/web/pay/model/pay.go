package model

type PayParams struct {
	Paytype     int         `json:"paytype" form:"paytype"`
	PrepayToken string      `json:"prepay_token" form:"prepay_token"`
	Notify      string      `json:"-"`
	Wexin       WxPayParams `json:"weixin"`
}

type WxPayParams struct {
	TradeType string `json:"trade_type"`
}
