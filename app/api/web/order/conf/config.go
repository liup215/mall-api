package conf

import (
	"mall/lib/database/orm"
	"os"
)

var Conf *Config

type Config struct {
	Orm     *orm.Config
	Http    *HttpConfig
	Service *ServiceConfig
}

type HttpConfig struct {
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
			Port: "8092",
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
			Host:         "rm-2ze6a2u0843v7lfk3.mysql.rds.aliyuncs.com",
			Port:         "3306",
			Username:     "malldev",
			Password:     "^7M6aZGK4XZe9#397n4v#Nwl",
			Charset:      "utf8",
			Database:     "mall-dev",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
		Http: &HttpConfig{
			Port: "80",
		},
		Service: &ServiceConfig{
			Finance: "http://byq-shop-finance-dev.emao.xin",
		},
	}
	return nil
}

func production() error {
	Conf = &Config{
		Orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "rm-2ze6a2u0843v7lfk3.mysql.rds.aliyuncs.com",
			Port:         "3306",
			Username:     "bianyiquan",
			Password:     "CTHFDc3Xejnf23Xe",
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
		Service: &ServiceConfig{
			Finance: "http://byq-shop-finance.emao.xin",
		},
	}
	return nil
}
