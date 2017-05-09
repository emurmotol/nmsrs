package unemployedstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type UnemployedStatus struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]UnemployedStatus, error) {
	unEmpStats := []UnemployedStatus{}

	if err := db.UnemployedStatuses.Find(nil).Sort("+name").All(&unEmpStats); err != nil {
		return nil, err
	}
	return unEmpStats, nil
}

func (unEmpStat *UnemployedStatus) Insert() (int, error) {
	if err := db.UnemployedStatuses.Insert(unEmpStat); err != nil {
		return 0, err
	}
	return unEmpStat.ID, nil
}

func FindByID(id int) (*UnemployedStatus, error) {
	var unEmpStat UnemployedStatus

	if err := db.UnemployedStatuses.Find(bson.M{"id": id}).One(&unEmpStat); err != nil {
		return &unEmpStat, err
	}
	return &unEmpStat, nil
}
