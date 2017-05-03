package unemployedstatus

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type UnemployedStatus struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]UnemployedStatus, error) {
	unEmpStats := []UnemployedStatus{}

	if err := db.UnemployedStatuses.Find(bson.M{}).Sort("+name").All(&unEmpStats); err != nil {
		return nil, err
	}
	return unEmpStats, nil
}

func (unEmpStat *UnemployedStatus) Insert() (string, error) {
	unEmpStat.ID = bson.NewObjectId()

	if err := db.UnemployedStatuses.Insert(unEmpStat); err != nil {
		return "", err
	}
	return unEmpStat.ID.Hex(), nil
}

func Find(id string) (*UnemployedStatus, error) {
	var unEmpStat UnemployedStatus

	if !bson.IsObjectIdHex(id) {
		return &unEmpStat, models.ErrInvalidObjectID
	}

	if err := db.UnemployedStatuses.FindId(bson.ObjectIdHex(id)).One(&unEmpStat); err != nil {
		return &unEmpStat, err
	}
	return &unEmpStat, nil
}
