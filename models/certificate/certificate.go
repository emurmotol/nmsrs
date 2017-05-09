package certificate

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Certificate struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Certificate, error) {
	certs := []Certificate{}

	if err := db.Certificates.Find(nil).Sort("+name").All(&certs); err != nil {
		return nil, err
	}
	return certs, nil
}

func (cert *Certificate) Insert() (string, error) {
	cert.ObjectID = bson.NewObjectId()

	if err := db.Certificates.Insert(cert); err != nil {
		return "", err
	}
	return cert.ObjectID.Hex(), nil
}

func FindByID(id string) (*Certificate, error) {
	var cert Certificate

	if !bson.IsObjectIdHex(id) {
		return &cert, models.ErrInvalidObjectID
	}

	if err := db.Certificates.FindId(bson.ObjectIdHex(id)).One(&cert); err != nil {
		return &cert, err
	}
	return &cert, nil
}
