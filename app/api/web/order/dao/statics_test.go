package dao

import (
	"mall/app/api/web/order/conf"
	"mall/app/api/web/order/model"
	"testing"
)

func TestSum(t *testing.T) {
	if err := conf.Init(); err != nil {
		t.Error(err.Error())
	}

	d := New(conf.Conf)

	money, err := d.OrderSumAndTotal(model.OrderStaticsParam{Status: -1})
	if err != nil {
		t.Error(err.Error())
	}

	t.Errorf("%v", money)

}
