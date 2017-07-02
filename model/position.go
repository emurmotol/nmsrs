package model

import (
	
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Position struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (position *Position) Create() *Position {
	if err := db.C("positions").Insert(position); err != nil {
		panic(err)
	}
	return position
}

func (position Position) Index(q string) []Position {
	positions := []Position{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"value": regex}

	if err := db.C("positions").Find(query).All(&positions); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return positions
}

func PositionById(id bson.ObjectId) *Position {
	position := new(Position)

	if err := db.C("positions").FindId(id).One(&position); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return position
}
