package nsq

import (
	"github.com/nsqio/go-nsq"
)

type Producer struct {
	prod   *nsq.Producer
	prefix string
}

func NewProducer(c *Config) *Producer {
	proddd, err := nsq.NewProducer(c.Host+":"+c.Port, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	return &Producer{
		prod:   proddd,
		prefix: c.Prefix,
	}
}

func (p *Producer) Publish(topic string, data []byte) error {
	return p.prod.Publish(p.prefix+topic, data)
}
