package position

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Position struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Position, error) {
	poss := []Position{}

	if err := db.Positions.Find(nil).Sort("+name").All(&poss); err != nil {
		return nil, err
	}
	return poss, nil
}

func (pos *Position) Insert() (int, error) {
	if err := db.Positions.Insert(pos); err != nil {
		return 0, err
	}
	return pos.ID, nil
}

func FindByID(id int) (*Position, error) {
	var pos Position

	if err := db.Positions.Find(bson.M{"id": id}).One(&pos); err != nil {
		return &pos, err
	}
	return &pos, nil
}

func Search(query interface{}) ([]Position, error) {
	poss := []Position{}

	if err := db.Positions.Find(query).Sort("+name").All(&poss); err != nil {
		return nil, err
	}
	return poss, nil
}
