package auth

import (
	memberApi "mall/app/service/main/member/api"

	"errors"
	"mall/lib/strings"
)

func (s *Service) Auth(uniacid int, mobile, pwd string) (*memberApi.MemberInfoReply, error) {
	record, err := s.mclient.MemberInfoByMobile(uniacid, mobile)
	if err != nil {
		return nil, err
	}

	if record.Pwd != strings.Md5(pwd+record.Salt) {
		return nil, errors.New("账号或密码错误")
	}

	return record, nil

}
