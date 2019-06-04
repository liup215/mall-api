package conf

import (
	"mall/lib/net/http/middleware/auth"
	"os"
)

var (
	Conf *Config
)

type Config struct {
	Port             string
	AuthConfig       *auth.Config
	MmbServerAddress string
}

func local() error {
	Conf = &Config{
		Port:             3000,
		MmbServerAddress: "",
	}
	return nil
}

func develop() error {
	Conf = &Config{
		Port: 80,
	}
	return nil
}

func production() error {
	Conf = &Config{
		Port: 80,
	}
	return nil
}

func Init() error {
	if os.Getenv("profile") == "production" {
		return production()
	}

	if os.Getenv("profile") == "development" {
		return develop()
	}

	return local()
}
