package civilstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type CivilStatus struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]CivilStatus, error) {
	civStats := []CivilStatus{}

	if err := db.CivilStatuses.Find(nil).All(&civStats); err != nil {
		return nil, err
	}
	return civStats, nil
}

func (civStat *CivilStatus) Insert() (string, error) {
	civStat.ObjectID = bson.NewObjectId()

	if err := db.CivilStatuses.Insert(civStat); err != nil {
		return "", err
	}
	return civStat.ObjectID.Hex(), nil
}

func FindByID(id string) (*CivilStatus, error) {
	var civStat CivilStatus

	if !bson.IsObjectIdHex(id) {
		return &civStat, models.ErrInvalidObjectID
	}

	if err := db.CivilStatuses.FindId(bson.ObjectIdHex(id)).One(&civStat); err != nil {
		return &civStat, err
	}
	return &civStat, nil
}
