package civilstatus

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type CivilStatus struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]CivilStatus, error) {
	civStats := []CivilStatus{}

	if err := db.CivilStatuses.Find(nil).All(&civStats); err != nil {
		return nil, err
	}
	return civStats, nil
}

func (civStat *CivilStatus) Insert() (int, error) {
	if err := db.CivilStatuses.Insert(civStat); err != nil {
		return 0, err
	}
	return civStat.ID, nil
}

func FindByID(id int) (*CivilStatus, error) {
	var civStat CivilStatus

	if err := db.CivilStatuses.Find(bson.M{"id": id}).One(&civStat); err != nil {
		return &civStat, err
	}
	return &civStat, nil
}
