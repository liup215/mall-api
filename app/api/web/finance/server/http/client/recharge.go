package client

import (
	"encoding/json"
	"errors"
	"mall/app/service/main/finance/model"
)

func (c *Client) PreRecharge(p *model.PreRechargeParam) (*model.PreRechargeResponse, error) {
	res, err := c.c.JSON("/recharge/logConfirm", p)
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Code    int                        `json:"code"`
		Message string                     `json:"message"`
		Data    *model.PreRechargeResponse `json:"data"`
	}

	var r Resp
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, errors.New(r.Message)
	}

	return r.Data, nil
}

func (c *Client) RechargeLogConfirm(p *model.RechargeLogConfirmParam) (*model.RechargeLogConfirmResponse, error) {

	res, err := c.c.RestfulPOST("/recharge/logConfirm", p)
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Code    int                               `json:"code"`
		Message string                            `json:"message"`
		Data    *model.RechargeLogConfirmResponse `json:"data"`
	}

	var r Resp
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, errors.New(r.Message)
	}

	return r.Data, nil
}
