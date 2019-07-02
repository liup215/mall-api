package nsq

import (
	"github.com/nsqio/go-nsq"
)

func NewConsumer(topic, channel string, handler nsq.Handler) (*nsq.Consumer, error) {
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	c.AddHandler(handler)
	return c, nil
}
