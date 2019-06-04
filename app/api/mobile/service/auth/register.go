package auth

func (s *Service) Register(uniacid int, mobile, pwd, pwdconfirm string) error {
	if uniacid == 0 {
		return errors.New("无效的应用id")
	}

	if mobile == "" {
		return errors.New("请输入正确的手机号码！")
	}

	if pwd == "" {
		return errors.New("请输入密码！")
	}

	return s.mclient.MemberCreate(uniacid, mobile, pwd)
}
