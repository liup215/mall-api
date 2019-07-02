package service

import (
	"mall/app/api/web/groups/conf"
	"mall/app/api/web/groups/model"
	"testing"
)

func TestActivityDetail(t *testing.T) {
	conf.Init()

	svr := New(conf.Conf)
	act, err := svr.QueryActivityOne(model.ActivityQuery{Id: "5d03733bd4bd6003034a6537", Status: 0})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Error(act)
	}
}
