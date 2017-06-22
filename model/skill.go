package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Skill struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func skillSeeder() {
	data := []string{
		"AUTO MECHANIC",
		"BEAUTICIAN",
		"CARPENTRY WORK",
		"COMPUTER LITERATE",
		"DOMESTIC CHORES",
		"DRIVER",
		"ELECTRICIAN",
		"EMBROIDERY",
		"GARDENING",
		"MASONRY",
		"PAINTER/ARTIST",
		"PAINTING JOBS",
		"PHOTOGRAPHY",
		"SEWING DRESSES",
		"STENOGRAPHY",
		"TAILORING",
	}

	for _, name := range data {
		skill := Skill{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		skill.Create()
	}
}

func (skill *Skill) Create() *Skill {
	db.C("skills").Insert(skill)
	defer db.Close()
	return skill
}

func (skill Skill) Index(q string) []Skill {
	skills := []Skill{}
	r := make(chan []Skill)
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	go func() {
		db.C("skills").Find(query).All(&skills)
		defer db.Close()
		r <- skills
	}()

	skills = <-r
	close(r)
	return skills
}
