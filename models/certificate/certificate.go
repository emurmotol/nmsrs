package certificate

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Certificate struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Certificate, error) {
	certs := []Certificate{}

	if err := db.Certificates.Find(bson.M{}).Sort("+name").All(&certs); err != nil {
		return nil, err
	}
	return certs, nil
}

func (cert *Certificate) Insert() (string, error) {
	cert.ID = bson.NewObjectId()

	if err := db.Certificates.Insert(cert); err != nil {
		return "", err
	}
	return cert.ID.Hex(), nil
}

func Find(id string) (*Certificate, error) {
	var cert Certificate

	if !bson.IsObjectIdHex(id) {
		return &cert, models.ErrInvalidObjectID
	}

	if err := db.Certificates.FindId(bson.ObjectIdHex(id)).One(&cert); err != nil {
		return &cert, err
	}
	return &cert, nil
}
