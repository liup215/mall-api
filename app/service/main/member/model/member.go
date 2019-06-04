package model

import (
	"mall/lib/time"
)

type EweiShopMember struct {
	Id                  int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid             int       `json:"uniacid,omitempty"`
	Uid                 int       `json:"uid,omitempty"`
	Groupid             int       `json:"groupid,omitempty"`
	Level               int       `json:"level,omitempty"`
	Agentid             int       `json:"agentid,omitempty"`
	Openid              string    `json:"openid,omitempty"`
	Realname            string    `json:"realname,omitempty"`
	Mobile              string    `json:"mobile,omitempty"`
	Pwd                 string    `json:"-"`
	Weixin              string    `json:"weixin,omitempty"`
	Content             string    `json:"content,omitempty"`
	Createtime          time.Time `json:"createtime,omitempty"`
	Agenttime           time.Time `json:"agenttime,omitempty"`
	Status              int       `json:"status,omitempty"`
	Isagent             int       `json:"isagent,omitempty"`
	Clickcount          int       `json:"clickcount,omitempty"`
	Agentlevel          int       `json:"agentlevel,omitempty"`
	Noticeset           string    `json:"noticeset,omitempty"`
	Nickname            string    `json:"nickname,omitempty"`
	Credit1             float64   `json:"credit1"`
	Credit2             float64   `json:"credit2"`
	Birthyear           string    `json:"birthyear,omitempty"`
	Birthmonth          string    `json:"birthmonth,omitempty"`
	Birthday            string    `json:"birthday,omitempty"`
	Gender              int       `json:"gender,omitempty"`
	Avatar              string    `json:"avatar,omitempty"`
	Province            string    `json:"province,omitempty"`
	City                string    `json:"city,omitempty"`
	Area                string    `json:"area,omitempty"`
	Childtime           time.Time `json:"childtime,omitempty"`
	Inviter             int       `json:"inviter,omitempty"`
	Agentnotupgrade     int       `json:"agentnotupgrade,omitempty"`
	Agentselectgoods    int       `json:"agentselectgoods,omitempty"`
	Agentblack          int       `json:"agentblack,omitempty"`
	Fixagentid          int       `json:"fixagentid,omitempty"`
	Diymemberid         int       `json:"diymemberid,omitempty"`
	Diymemberfields     string    `json:"diymemberfields,omitempty"`
	Diymemberdata       string    `json:"diymemberdata,omitempty"`
	Diymemberdataid     int       `json:"diymemberdataid,omitempty"`
	Diycommissionid     int       `json:"diycommissionid,omitempty"`
	Diycommissionfields string    `json:"diycommissionfields,omitempty"`
	Diycommissiondata   string    `json:"diycommissiondata,omitempty"`
	Diycommissiondataid int       `json:"diycommissiondataid,omitempty"`
	Isblack             int       `json:"isblack,omitempty"`
	Username            string    `json:"username,omitempty"`
	CommissionTotal     float64   `json:"commission_total,omitempty"`
	Endtime2            time.Time `json:"endtime2,omitempty"`
	Ispartner           int       `json:"ispartner,omitempty"`
	Partnertime         time.Time `json:"partnertime,omitempty"`
	Partnerstatus       int       `json:"partnerstatus,omitempty"`
	Partnerblack        int       `json:"partnerblack,omitempty"`
	Partnerlevel        int       `json:"partnerlevel,omitempty"`
	Partnernotupgrade   int       `json:"partnernotupgrade,omitempty"`
	Diyglobonusid       int       `json:"diyglobonusid,omitempty"`
	Diyglobonusdata     string    `json:"diyglobonusdata,omitempty"`
	Diyglobonusfields   string    `json:"diyglobonusfields,omitempty"`
	Isaagent            int       `json:"isaagent,omitempty"`
	Aagentlevel         int       `json:"aagentlevel,omitempty"`
	Aagenttime          time.Time `json:"aagenttime,omitempty"`
	Aagentstatus        int       `json:"aagentstatus,omitempty"`
	Aagentblack         int       `json:"aagentblack,omitempty"`
	Aagentnotupgrade    int       `json:"aagentnotupgrade,omitempty"`
	Aagenttype          int       `json:"aagenttype,omitempty"`
	Aagentprovinces     string    `json:"aagentprovinces,omitempty"`
	Aagentcitys         string    `json:"aagentcitys,omitempty"`
	Aagentareas         string    `json:"aagentareas,omitempty"`
	Diyaagentid         int       `json:"diyaagentid,omitempty"`
	Diyaagentdata       string    `json:"diyaagentdata,omitempty"`
	Diyaagentfields     string    `json:"diyaagentfields,omitempty"`
	Salt                string    `json:"-"`
	Mobileverify        int       `json:"mobileverify,omitempty"`
	Mobileuser          int       `json:"mobileuser,omitempty"`
	CarrierMobile       string    `json:"carrier_mobile,omitempty"`
	Isauthor            int       `json:"isauthor,omitempty"`
	Authortime          time.Time `json:"authortime,omitempty"`
	Authorstatus        int       `json:"authorstatus,omitempty"`
	Authorblack         int       `json:"authorblack,omitempty"`
	Authorlevel         int       `json:"authorlevel,omitempty"`
	Authornotupgrade    int       `json:"authornotupgrade,omitempty"`
	Diyauthorid         int       `json:"diyauthorid,omitempty"`
	Diyauthordata       string    `json:"diyauthordata,omitempty"`
	Diyauthorfields     string    `json:"diyauthorfields,omitempty"`
	Authorid            int       `json:"authorid,omitempty"`
	Comefrom            string    `json:"comefrom,omitempty"`
	OpenidQq            string    `json:"openid_qq,omitempty"`
	OpenidWx            string    `json:"openid_wx,omitempty"`
	Diymaxcredit        int       `json:"diymaxcredit,omitempty"`
	Maxcredit           int       `json:"maxcredit,omitempty"`
	Datavalue           string    `json:"datavalue,omitempty"`
	OpenidWa            string    `json:"openid_wa,omitempty"`
	NicknameWechat      string    `json:"nickname_wechat,omitempty"`
	AvatarWechat        string    `json:"avatar_wechat,omitempty"`
	Updateaddress       int       `json:"updateaddress,omitempty"`
	Membercardid        string    `json:"membercardid,omitempty"`
	Membercardcode      string    `json:"membercardcode,omitempty"`
	Membershipnumber    string    `json:"membershipnumber,omitempty"`
	Membercardactive    int       `json:"membercardactive,omitempty"`
	Commission          float64   `json:"commission,omitempty"`
	CommissionPay       float64   `json:"commission_pay,omitempty"`
	Idnumber            string    `json:"idnumber,omitempty"`
	Wxcardupdatetime    time.Time `json:"wxcardupdatetime,omitempty"`
}

type MemberQuery struct {
	Id           int    `json:"id" form:"id"`
	Uniacid      int    `json:"uniacid" form:"uniacid"`
	Openid       string `json:"openid" form:"openid"`
	Mobile       string `json:"mobile" form:"mobile"`
	Mobileverify int    `json:"mobileverify" form:"mobileverify"`
}
