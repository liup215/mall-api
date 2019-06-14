package redis

import (
	"testing"
)

func TestRedis(t *testing.T) {
	c := &Config{
		Host: "www.emao.xin",
		Port: "6379",
		DB:   5,
	}

	redis := New(c)
	err := redis.SET("a", "b", 0)
	if err != nil {
		t.Error(err.Error())
	}
	v, err := redis.GET("a")
	if err != nil {
		t.Error(err.Error())
	}

	if v != "b" {
		t.Errorf("error, expect %s, got %s", "b", v)
	}
}
