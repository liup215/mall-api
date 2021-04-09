package model

import (
	"mall/lib/time"
)

type UserView struct {
	Uid            int       `json:"uid,omitempty"`
	OwnerUid       int       `json:"owner_uid,omitempty"`
	Groupid        int       `json:"groupid,omitempty"`
	FounderGroupid int       `json:"founder_groupid,omitempty"`
	Username       string    `json:"username,omitempty"`
	Password       string    `json:"-,omitempty"`
	Salt           string    `json:"-,omitempty"`
	Type           int       `json:"type,omitempty"`
	Status         int       `json:"status"`
	Joindate       time.Time `json:"joindate,omitempty"`
	Joinip         string    `json:"joinip,omitempty"`
	Lastvisit      int       `json:"lastvisit,omitempty"`
	Lastip         string    `json:"lastip,omitempty"`
	Remark         string    `json:"remark,omitempty"`
	Starttime      time.Time `json:"starttime,omitempty"`
	Endtime        time.Time `json:"endtime,omitempty"`
	Uniacid        int       `json:"uniacid,omitempty"`
	Tid            int       `json:"tid,omitempty"`
	Schoolid       int       `json:"schoolid,omitempty"`
	RegisterType   int       `json:"register_type,omitempty"`
	Openid         string    `json:"openid,omitempty"`
}
