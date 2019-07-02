package conf

import (
	"mall/lib/database/nosql/mongo"
	"mall/lib/database/orm"
	"mall/lib/database/redis"
	"mall/lib/net/http/middleware/auth"
	"os"
)

var Conf *Config

type Config struct {
	Orm   *orm.Config
	Mongo *mongo.Config
	Redis *redis.Config
	Http  *HttpConfig
	Grpc  *GrpcConfig
	Auth  *auth.Config
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
		Redis: &redis.Config{
			Host: "localhost",
			Port: "6379",
		},
		Http: &HttpConfig{
			Port: "8094",
		},
		Grpc: &GrpcConfig{
			Port: "8094",
		},
		Mongo: &mongo.Config{
			Host:     "localhost",
			Port:     "27017",
			Username: "root",
			Password: "123456",
			Db:       "mall",
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
		Mongo: &mongo.Config{
			Host:     "172.17.222.99",
			Port:     "27017",
			Username: "root",
			Password: "123456",
			Db:       "mall-dev",
		},
		Redis: &redis.Config{
			Host: "172.17.222.99",
			Port: "6379",
			DB:   5,
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
		Mongo: &mongo.Config{
			Host:     "172.17.222.99",
			Port:     "27017",
			Username: "root",
			Password: "123456",
			Db:       "mall",
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
