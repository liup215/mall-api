package model

type AccountWechats struct {
	Acid             int    `gorm:"primary_key;AUTO_INCREMENT" json:"acid"`
	Uniacid          int    `json:"uniacid"`
	Token            string `json:"token"`
	Encodingaeskey   string `json:"encodingaeskey"`
	Level            int    `json:"level"`
	Name             string `json:"name"`
	Account          string `json:"account"`
	Original         string `json:"original"`
	Signature        string `json:"signature"`
	Country          string `json:"country"`
	Province         string `json:"province"`
	City             string `json:"city"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Lastupdate       int    `json:"lastupdate"`
	Key              string `json:"key"`
	Secret           string `json:"secret"`
	Styleid          int    `json:"styleid"`
	Subscribeurl     string `json:"subscribeurl"`
	AuthRefreshToken string `json:"auth_refresh_token"`
	AccessToken      string `json:"access_token"`
}

type AccountWechatsQuery struct {
	Acid    int `json:"acid" form:"acid"`
	Uniacid int `json:"uniacid" form:"uniacid"`
}

type WechatOauthParamsReply struct {
	Appid string `json:"appid"`
	State string `json:"state"`
}

type WechatUserInfoParams struct {
	Uniacid int    `json:"uniacid" form:"uniacid"`
	Code    string `json:"code" form:"code"`
	State   string `json:"state" form:"state"`
}
