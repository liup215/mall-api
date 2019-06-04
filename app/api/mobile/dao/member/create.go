package member

import (
	"mall/app/service/main/member/api"
)

func (d *Dao) Create(uniacid int, mobile, pwd string) (*api.MemberCreateReply, error) {
	return d.mclit.MemberCreate(uniacid, mobile, pwd)
}
