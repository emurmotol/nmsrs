package employmentstatus

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type EmploymentStatus struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]EmploymentStatus, error) {
	empStats := []EmploymentStatus{}

	if err := db.EmploymentStatuses.Find(bson.M{}).All(&empStats); err != nil {
		return nil, err
	}
	return empStats, nil
}

func (empStat *EmploymentStatus) Insert() (string, error) {
	empStat.ID = bson.NewObjectId()

	if err := db.EmploymentStatuses.Insert(empStat); err != nil {
		return "", err
	}
	return empStat.ID.Hex(), nil
}

func Find(id string) (*EmploymentStatus, error) {
	var empStat EmploymentStatus

	if !bson.IsObjectIdHex(id) {
		return &empStat, models.ErrInvalidObjectID
	}

	if err := db.EmploymentStatuses.FindId(bson.ObjectIdHex(id)).One(&empStat); err != nil {
		return &empStat, err
	}
	return &empStat, nil
}
