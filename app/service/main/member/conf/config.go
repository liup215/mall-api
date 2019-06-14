package conf

import (
	"mall/lib/database/orm"
	"mall/lib/net/http/middleware/auth"
	"os"
)

var Conf *Config

type Config struct {
	Auth    *auth.Config
	Orm     *orm.Config
	Http    *HttpConfig
	Grpc    *GrpcConfig
	Service *ServiceConfig
}

type HttpConfig struct {
	Port string
}

type GrpcConfig struct {
	Port string
}

type ServiceConfig struct {
	Member  string
	Finance string
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
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf8",
			Database:     "emao_shop",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "8090",
		},
		Grpc: &GrpcConfig{
			Port: "8090",
		},
		Service: &ServiceConfig{
			Finance: "http://localhost:8091",
		},
	}
	return nil
}
func development() error {
	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.222.99",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf8",
			Database:     "bianyiquan_dev",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "80",
		},
		Grpc: &GrpcConfig{
			Port: "80",
		},
		Service: &ServiceConfig{
			Finance: "http://byq-shop-finance.emao.xin",
		},
	}
	return nil
}

func production() error {
	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.222.99",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf8",
			Database:     "bianyiquan",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "80",
		},
		Grpc: &GrpcConfig{
			Port: "80",
		},
		Service: &ServiceConfig{
			Finance: "http://byq-shop-finance.emao.xin",
		},
	}
	return nil
}
