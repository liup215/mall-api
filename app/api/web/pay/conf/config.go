package conf

import (
	"mall/lib/net/http/middleware/auth"
	"mall/lib/net/rpc"
	"os"
)

var Conf *Config

type Config struct {
	PaymentClient *rpc.ClientConfig
	Auth          *auth.Config
	Http          *HttpConfig
}

type HttpConfig struct {
	Port string
}

func Init() error {

	if os.Getenv("profile") == "production" {
		return production()
	}

	if os.Getenv("profile") == "development" {
		return development()
	}

	return local()

}

func local() error {
	Conf = &Config{
		PaymentClient: &rpc.ClientConfig{
			Host: "localhost",
			Port: "8087",
		},
		Http: &HttpConfig{
			Port: "8097",
		},
	}
	return nil
}
func development() error {
	Conf = &Config{
		PaymentClient: &rpc.ClientConfig{
			Host: "byq-shop-payment-dev.grpc.emao.xin",
			Port: "80",
		},
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}

func production() error {
	Conf = &Config{
		PaymentClient: &rpc.ClientConfig{
			Host: "byq-shop-payment.grpc.emao.xin",
			Port: "80",
		},
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}
