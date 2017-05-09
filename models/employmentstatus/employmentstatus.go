package employmentstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type EmploymentStatus struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]EmploymentStatus, error) {
	empStats := []EmploymentStatus{}

	if err := db.EmploymentStatuses.Find(nil).All(&empStats); err != nil {
		return nil, err
	}
	return empStats, nil
}

func (empStat *EmploymentStatus) Insert() (int, error) {
	if err := db.EmploymentStatuses.Insert(empStat); err != nil {
		return 0, err
	}
	return empStat.ID, nil
}

func FindByID(id int) (*EmploymentStatus, error) {
	var empStat EmploymentStatus

	if err := db.EmploymentStatuses.Find(bson.M{"id": id}).One(&empStat); err != nil {
		return &empStat, err
	}
	return &empStat, nil
}
