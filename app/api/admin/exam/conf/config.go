package conf

import (
	"mall/lib/net/http/middleware/auth"
	"os"
)

var Conf *Config

type Config struct {
	Auth *auth.Config
	Http *HttpConfig
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
		Http: &HttpConfig{
			Port: "8090",
		},
	}
	return nil
}
func development() error {
	Conf = &Config{
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}

func production() error {
	Conf = &Config{
		Http: &HttpConfig{
			Port: "80",
		},
	}
	return nil
}
