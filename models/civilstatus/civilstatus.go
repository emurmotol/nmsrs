package civilstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type CivilStatus struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]CivilStatus, error) {
	civStats := []CivilStatus{}

	if err := db.CivilStatuses.Find(bson.M{}).All(&civStats); err != nil {
		return nil, err
	}
	return civStats, nil
}

func (civStat *CivilStatus) Insert() (string, error) {
	civStat.ID = bson.NewObjectId()

	if err := db.CivilStatuses.Insert(civStat); err != nil {
		return "", err
	}
	return civStat.ID.Hex(), nil
}

func Find(id string) (*CivilStatus, error) {
	var civStat CivilStatus

	if !bson.IsObjectIdHex(id) {
		return &civStat, models.ErrInvalidObjectID
	}

	if err := db.CivilStatuses.FindId(bson.ObjectIdHex(id)).One(&civStat); err != nil {
		return &civStat, err
	}
	return &civStat, nil
}
