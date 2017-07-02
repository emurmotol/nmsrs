package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"github.com/emurmotol/nmsrs/env"
)

var (
	Name    string
	Session *mgo.Session
)

func init() {
	Name, _ = env.Conf.String("pkg.mgo.name")

	Session = func() *mgo.Session {
		host, _ := env.Conf.String("pkg.mgo.host")
		port, _ := env.Conf.Int("pkg.mgo.port")
		s, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", host, port))

		if err != nil {
			panic(err)
		}
		return s
	}()
}

func Con() *mgo.Database {
	return Session.DB(Name)
}
