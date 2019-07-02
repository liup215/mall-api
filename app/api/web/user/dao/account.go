package dao

import (
	"errors"
	"mall/app/api/web/member/model"
)

func (d *Dao) UpdatePwd(p *model.UpdatePwdParam) error {
	if p.Uniacid == 0 {
		return errors.New("无效的应用ID")
	}
	if p.Openid == "" {
		return errors.New("无效的Openid")
	}

	if p.Newpwd == "" {
		return errors.New("请输入新密码")
	}

	return d.orm.Model(&model.EweiShopMember{}).Where(&model.EweiShopMember{Uniacid: p.Uniacid, Openid: p.Openid}).Update(map[string]interface{}{
		"pwd": p.Newpwd,
	}).Error
}
