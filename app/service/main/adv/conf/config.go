package conf

import (
	"mall/lib/database/orm"
	"os"
)

var conf *Config

type Config struct {
	orm *orm.Config
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

	conf := &Config{
		orm: &orm.Config{
			Dialect:      "mysql",
			Host:         "localhost",
			Port:         "3306",
			Username:     "root",
			Password:     "123456",
			Charset:      "utf-8",
			Database:     "emao-shop",
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "ims_",
		},
	}
	return nil
}
func development() error {
	return nil
}

func production() error {
	return nil
}
