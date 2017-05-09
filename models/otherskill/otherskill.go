package otherskill

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type OtherSkill struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]OtherSkill, error) {
	oskills := []OtherSkill{}

	if err := db.OtherSkills.Find(nil).Sort("+name").All(&oskills); err != nil {
		return nil, err
	}
	return oskills, nil
}

func (oskill *OtherSkill) Insert() (int, error) {
	if err := db.OtherSkills.Insert(oskill); err != nil {
		return 0, err
	}
	return oskill.ID, nil
}

func FindByID(id int) (*OtherSkill, error) {
	var oskill OtherSkill

	if err := db.OtherSkills.Find(bson.M{"id": id}).One(&oskill); err != nil {
		return &oskill, err
	}
	return &oskill, nil
}

func Search(query interface{}) ([]OtherSkill, error) {
	oskills := []OtherSkill{}

	if err := db.OtherSkills.Find(query).Sort("+name").All(&oskills); err != nil {
		return nil, err
	}
	return oskills, nil
}
