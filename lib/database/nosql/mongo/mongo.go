package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Db       string
}

func New(c *Config) *Mongo {
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", c.Username, c.Password, c.Host, c.Port, c.Db)
	fmt.Println("mongo:", url)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	return &Mongo{
		session: session,
		db:      c.Db,
	}
}

type Mongo struct {
	session *mgo.Session
	db      string
}

func (mgoo *Mongo) Model(collection string) *Model {
	ss := mgoo.session.Copy()
	c := ss.DB(mgoo.db).C(collection)

	return &Model{
		session: ss,
		C:       c,
	}

}

type Model struct {
	session *mgo.Session
	C       *mgo.Collection
}

func (model *Model) Close() {
	model.session.Close()
}
