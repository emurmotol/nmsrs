package db

import (
	"fmt"

	"github.com/zneyrl/nmsrs/env"
	mgo "gopkg.in/mgo.v2"
)

var (
	DB          *mgo.Database
	Users       *mgo.Collection
	Registrants *mgo.Collection
)

func init() {
	s, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", env.DBHost, env.DBPort))

	if err != nil {
		panic(err)
	}

	if err := s.Ping(); err != nil {
		panic(err)
	}
	DB = s.DB(env.DBName)
	Users = DB.C("users")
	Registrants = DB.C("registrants")
}
