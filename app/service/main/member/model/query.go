package model

type MemberQuery struct {
	Id      int    `json:"id" from:"id"`
	Uniacid int    `json:"uniacid" from:"uniacid"`
	Openid  string `json:"openid" from:"openid"`
	Mobile  string `json:"mobile" from:"mobile"`
}
