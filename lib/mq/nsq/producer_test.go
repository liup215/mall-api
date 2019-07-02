package nsq

import (
	"testing"
)

func TestProducer(t *testing.T) {
	p := NewProducer(&Config{
		Host: "www.emao.xin",
		Port: "4150",
	})

	err := p.Publish("test", []byte("test"))
	if err != nil {
		t.Error(err.Error())
	}
}
