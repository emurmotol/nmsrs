package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Certificate struct {
	Id    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
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

func CertificateById(id bson.ObjectId) *Certificate {
	certificate := new(Certificate)

	if err := db.C("certificates").FindId(id).One(&certificate); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return certificate
}
