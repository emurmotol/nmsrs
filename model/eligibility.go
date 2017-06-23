package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func eligibilitySeeder() {
	data := []string{
		"CAREER EXECUTIVE OFFICER ELIGIBILITY",
		"CAREER EXECUTIVE SERVICE OFFICER",
		"CAREER SERVICE EXECUTIVE ELIGIBILITY",
		"CAREER SERVICE PROFESSIONAL",
		"CAREER SERVICE SUB - PROFESSIONAL",
		"DATA ENCODER",
		"FIRE OFFICER 2",
		"FORESTRY EXTENSION SERVICE",
		"POLICE OFFICER 1",
		"R.A. 1080",
		"SOIL TECHNOLOGIST",
		"STENOGRAPHER",
	}

	for _, name := range data {
		eligibility := Eligibility{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		eligibility.Create()
	}
}

type Eligibility struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func (eligibility *Eligibility) Create() *Eligibility {
	if err := db.C("eligibilities").Insert(eligibility); err != nil {
		panic(err)
	}
	defer db.Close()
	return eligibility
}

func (eligibility Eligibility) Index(q string) []Eligibility {
	eligibilities := []Eligibility{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("eligibilities").Find(query).All(&eligibilities); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return eligibilities
}
