package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func certificateSeeder() {
	data, err := ioutil.ReadFile("import/certificates.json")

	if err != nil {
		panic(err)
	}
	certificates := []Certificate{}

	if err := json.Unmarshal(data, &certificates); err != nil {
		panic(err)
	}
	log.Println("certificateSeeder: todo")
}

type Certificate struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (certificate *Certificate) Create() *Certificate {
	if err := db.C("certificates").Insert(certificate); err != nil {
		panic(err)
	}
	defer db.Close()
	return certificate
}

func (certificate Certificate) Index(q string) []Certificate {
	certificates := []Certificate{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("certificates").Find(query).All(&certificates); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return certificates
}
