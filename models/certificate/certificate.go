package certificate

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Certificate struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Certificate, error) {
	certs := []Certificate{}

	if err := db.Certificates.Find(nil).Sort("+name").All(&certs); err != nil {
		return nil, err
	}
	return certs, nil
}

func (cert *Certificate) Insert() (int, error) {
	if err := db.Certificates.Insert(cert); err != nil {
		return 0, err
	}
	return cert.ID, nil
}

func FindByID(id int) (*Certificate, error) {
	var cert Certificate

	if err := db.Certificates.Find(bson.M{"id": id}).One(&cert); err != nil {
		return &cert, err
	}
	return &cert, nil
}

func Search(query interface{}) ([]Certificate, error) {
	certs := []Certificate{}

	if err := db.Certificates.Find(query).Sort("+name").All(&certs); err != nil {
		return nil, err
	}
	return certs, nil
}
