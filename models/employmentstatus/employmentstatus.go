package employmentstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type EmploymentStatus struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]EmploymentStatus, error) {
	empStats := []EmploymentStatus{}

	if err := db.EmploymentStatuses.Find(nil).All(&empStats); err != nil {
		return nil, err
	}
	return empStats, nil
}

func (empStat *EmploymentStatus) Insert() (string, error) {
	empStat.ObjectID = bson.NewObjectId()

	if err := db.EmploymentStatuses.Insert(empStat); err != nil {
		return "", err
	}
	return empStat.ObjectID.Hex(), nil
}

func FindByID(id string) (*EmploymentStatus, error) {
	var empStat EmploymentStatus

	if !bson.IsObjectIdHex(id) {
		return &empStat, models.ErrInvalidObjectID
	}

	if err := db.EmploymentStatuses.FindId(bson.ObjectIdHex(id)).One(&empStat); err != nil {
		return &empStat, err
	}
	return &empStat, nil
}
