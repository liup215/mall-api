package conf

import (
	"mall/lib/database/orm"
	"mall/lib/net/http/middleware/auth"
	"os"
)

var Conf *Config

type Config struct {
	AuthKey string
	Orm     *orm.Config
	Http    *HttpConfig
	Grpc    *GrpcConfig
	Auth    *auth.Config
}

type HttpConfig struct {
	Port string
}

type GrpcConfig struct {
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
		AuthKey: "35eb5e04",
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf8",
			Database:     "emao_exam",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "8093",
		},
		Grpc: &GrpcConfig{
			Port: "8093",
		},
	}
	return nil
}
func development() error {
	Conf = &Config{
		AuthKey: "35eb5e04",
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.0.1",
			Port:         "3306",
			Username:     "root",
			Password:     "admin",
			Charset:      "utf8",
			Database:     "emao_exam_dev",
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
	}
	return nil
}

func production() error {
	Conf = &Config{
		AuthKey: "35eb5e04",
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "172.17.o.1",
			Port:         "3306",
			Username:     "root",
			Password:     "admin",
			Charset:      "utf8",
			Database:     "emao_exam",
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
	}
	return nil
}
