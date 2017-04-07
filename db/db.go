package db

import (
	"log"

	couchdb "github.com/rhinoman/couchdb-go"
	"github.com/zneyrl/nmsrs-lookup/env"
)

var Con *couchdb.Database

func init() {
	c, err := couchdb.NewConnection(env.DBHost, env.DBPort, env.DBTimeout)

	if err != nil {
		log.Fatal(err)
	}
	auth := couchdb.BasicAuth{env.DBUser, env.DBPassword}
	Con = c.SelectDB(env.DBName, &auth)
}
