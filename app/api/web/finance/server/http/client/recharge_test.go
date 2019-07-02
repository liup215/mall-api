package client

import (
	"testing"
)

func TestRechargeConfirm(t *testing.T) {
	host := "http://localhost:8091"
	c := New(host)

	res, err := c.RechargeConfirm("RC20190518123615155223", 1)

	if err != nil {
		t.Error(err.Error())
	}

	t.Error(res)

}
