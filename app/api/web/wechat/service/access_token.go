package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mall/app/api/web/wechat/model"
	"net/http"
)

func (s *Service) GetWechatService(uniacid int) (*WechatService, error) {
	account, err := s.d.QueryAccountWechatsOne(model.AccountWechatsQuery{Uniacid: uniacid})
	if err != nil {
		return nil, err
	}
	return &WechatService{
		acid:      account.Acid,
		appid:     account.Key,
		appsecret: account.Secret,
		svr:       s,
	}, nil
}

type WechatService struct {
	acid      int
	appid     string
	appsecret string
	svr       *Service
}

func (s *WechatService) Token() (string, error) {
	cachekey := fmt.Sprintf("accesstoken:%v", s.acid)
	token, err := s.svr.d.GetCache(cachekey)
	if err == nil {
		return token, err
	}

	return s.RefreshToken("")
}

func (s *WechatService) RefreshToken(currentToken string) (string, error) {
	cachekey := fmt.Sprintf("accesstoken:%v", s.acid)
	if s.appid == "" || s.appsecret == "" {
		return "", errors.New("未填写公众号的 appid 或 appsecret！")
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", s.appid, s.appsecret)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type AccessTokenResponse struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	var accessToken AccessTokenResponse

	if err := json.Unmarshal(body, &accessToken); err != nil {
		return "", err
	}

	if accessToken.ErrCode != 0 {
		return "", errors.New(accessToken.ErrMsg)
	}

	s.svr.d.SetCache(cachekey, accessToken.AccessToken, accessToken.ExpiresIn)
	return accessToken.AccessToken, nil
}

func (srv *WechatService) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}
