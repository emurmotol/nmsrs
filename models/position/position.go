package position

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Position struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Position, error) {
	poss := []Position{}

	if err := db.Positions.Find(bson.M{}).Sort("+name").All(&poss); err != nil {
		return nil, err
	}
	return poss, nil
}

func (pos *Position) Insert() (string, error) {
	pos.ID = bson.NewObjectId()

	if err := db.Positions.Insert(pos); err != nil {
		return "", err
	}
	return pos.ID.Hex(), nil
}

func Find(id string) (*Position, error) {
	var pos Position

	if !bson.IsObjectIdHex(id) {
		return &pos, models.ErrInvalidObjectID
	}

	if err := db.Positions.FindId(bson.ObjectIdHex(id)).One(&pos); err != nil {
		return &pos, err
	}
	return &pos, nil
}
