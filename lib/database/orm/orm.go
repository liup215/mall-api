package orm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Dialect      string
	Host         string
	Port         string
	Username     string
	Password     string
	Database     string
	Charset      string
	SQLLog       bool
	MaxIdleConns int
	MaxOpenConns int
	Prefix       string
}

func New(conf *Config) *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)
	fmt.Println(url)
	db, err := gorm.Open(conf.Dialect, url)
	if err != nil {
		panic(err)
	}

	if conf.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.Prefix + defaultTableName
	}
	return db
}
