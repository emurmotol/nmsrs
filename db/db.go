package db

import (
	"fmt"
	"log"

	"github.com/zneyrl/nmsrs-lookup/env"
	mgo "gopkg.in/mgo.v2"
)

var (
	DB    *mgo.Database
	Users *mgo.Collection
)

func init() {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", env.DBHost, env.DBPort))

	if err != nil {
		log.Fatal(err)
	}

	if err := session.Ping(); err != nil {
		log.Fatal(err)
	}
	DB = session.DB(env.DBName)
	Users = DB.C("users")
}
