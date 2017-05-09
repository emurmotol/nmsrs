package unemployedstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type UnemployedStatus struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]UnemployedStatus, error) {
	unEmpStats := []UnemployedStatus{}

	if err := db.UnemployedStatuses.Find(nil).Sort("+name").All(&unEmpStats); err != nil {
		return nil, err
	}
	return unEmpStats, nil
}

func (unEmpStat *UnemployedStatus) Insert() (string, error) {
	unEmpStat.ObjectID = bson.NewObjectId()

	if err := db.UnemployedStatuses.Insert(unEmpStat); err != nil {
		return "", err
	}
	return unEmpStat.ObjectID.Hex(), nil
}

func FindByID(id string) (*UnemployedStatus, error) {
	var unEmpStat UnemployedStatus

	if !bson.IsObjectIdHex(id) {
		return &unEmpStat, models.ErrInvalidObjectID
	}

	if err := db.UnemployedStatuses.FindId(bson.ObjectIdHex(id)).One(&unEmpStat); err != nil {
		return &unEmpStat, err
	}
	return &unEmpStat, nil
}
