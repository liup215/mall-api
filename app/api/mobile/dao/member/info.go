package member

import (
	memberApi "mall/app/service/main/member/api"
)

func (d *Dao) MemberInfoByMobile(uniacid int, mobile string) (*memberApi.MemberInfoReply, error) {
	return d.mclit.MemberInfoByMobile(uniacid, mobile)
}
