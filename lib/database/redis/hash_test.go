package redis

import (
	"testing"
)

func TestHMSET(t *testing.T) {
	redis := New(&Config{
		Host: "www.emao.xin",
		Port: "6379",
		DB:   5,
	})

	type PrepayOrder struct {
		Orderno   string  `redis:"orderno"`
		Module    string  `redis:"module"`
		Uniacid   int     `redis:"uniacid"`
		Openid    string  `redis:"openid"`
		Fee       float64 `redis:"fee"`
		Body      string  `redis:"body"`
		CreateIp  string  `redis:"create_ip"`
		NotifyURL string  `redis:"notify_url"`
	}
	var o PrepayOrder
	err := redis.HGETALL("0d6f9148096104f5a85a7cee90d99a1d", &o)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Error(o)
	}
}
