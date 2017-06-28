package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func courseSeeder() {
	data, err := ioutil.ReadFile("import/courses.json")

	if err != nil {
		panic(err)
	}
	courses := []Course{}

	if err := json.Unmarshal(data, &courses); err != nil {
		panic(err)
	}
	log.Println("courseSeeder: todo")
}

type Course struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (course *Course) Create() *Course {
	if err := db.C("courses").Insert(course); err != nil {
		panic(err)
	}
	defer db.Close()
	return course
}

func (course Course) Index(q string) []Course {
	courses := []Course{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("courses").Find(query).All(&courses); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return courses
}
