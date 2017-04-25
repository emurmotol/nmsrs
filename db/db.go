package db

import (
	"fmt"
	"log"

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
		log.Fatal(err)
	}

	if err := s.Ping(); err != nil {
		log.Fatal(err)
	}
	DB = s.DB(env.DBName)
	Users = DB.C("users")
	Registrants = DB.C("registrants")
}
