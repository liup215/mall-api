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

}
