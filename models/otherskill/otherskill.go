package otherskill

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type OtherSkill struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]OtherSkill, error) {
	oskills := []OtherSkill{}

	if err := db.OtherSkills.Find(bson.M{}).Sort("+name").All(&oskills); err != nil {
		return nil, err
	}
	return oskills, nil
}

func (oskill *OtherSkill) Insert() (string, error) {
	oskill.ID = bson.NewObjectId()

	if err := db.OtherSkills.Insert(oskill); err != nil {
		return "", err
	}
	return oskill.ID.Hex(), nil
}

func Find(id string) (*OtherSkill, error) {
	var oskill OtherSkill

	if !bson.IsObjectIdHex(id) {
		return &oskill, models.ErrInvalidObjectID
	}

	if err := db.OtherSkills.FindId(bson.ObjectIdHex(id)).One(&oskill); err != nil {
		return &oskill, err
	}
	return &oskill, nil
}
