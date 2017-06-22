package db

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"

	"github.com/emurmotol/nmsrs/env"
)

var Name string

func init() {
	Name, _ = env.Conf.String("pkg.mgo.name")
}

func C(name string) *mgo.Collection {
	return Session().DB(Name).C(name)
}

func Close() {
	Session().Close()
}

func Session() *mgo.Session {
	host, _ := env.Conf.String("pkg.mgo.host")
	port, _ := env.Conf.Int("pkg.mgo.port")
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", host, port))

	if err != nil {
		panic(err)
	}
	return session
}
