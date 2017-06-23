package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func otherSkillSeeder() {
	data := []string{
		"ANALYZING",
		"ANTICIPATING",
		"ASSEMBLING",
		"BUILDING",
		"CHECKING",
		"COMPARING",
		"COMPILING",
		"COMPUTING",
		"COORDINATING",
		"COPYING",
		"CREATING/INVENTING",
		"DISCOVERING",
		"DIVERTING",
		"DRIVING/STEERING",
		"ENCOURAGING",
		"EXPRESSING",
		"FEEDING/LOADING",
		"HELPING",
		"IMPLEMENTING",
		"INNOVATING",
		"INSPECTING",
		"INSTRUCTING",
		"INTERPRETING",
		"LEADING",
		"MACHINE WORK",
		"MANIPULATING",
		"MATERIALS HANDLING",
		"MOTIVATING",
		"NEGOTIATING",
		"OPERATING/CONTROLLING",
		"ORGANIZING",
		"PERSUADING",
		"PLANNING",
		"POSTING",
		"PRECISION WORKING",
		"PREDICTING",
		"PRODUCING",
		"PROMOTING",
		"RECORDING",
		"REPAIRING/ADJUSTING",
		"RESEARCHING",
		"SELLING",
		"SERVING",
		"SETTING - UP",
		"SETTING - UP/RESTORING",
		"SPEAKING",
		"SPECULATING",
		"SYNTHESIZING IDEA",
		"TABULATING",
		"TEACHING",
		"TESTING",
		"THEORIZING",
		"WAREHOUSING",
	}

	for _, name := range data {
		otherSkill := OtherSkill{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		otherSkill.Create()
	}
}

type OtherSkill struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func (otherSkill *OtherSkill) Create() *OtherSkill {
	if err := db.C("otherSkills").Insert(otherSkill); err != nil {
		panic(err)
	}
	defer db.Close()
	return otherSkill
}

func (otherSkill OtherSkill) Index(q string) []OtherSkill {
	otherSkills := []OtherSkill{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("otherSkills").Find(query).All(&otherSkills); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return otherSkills
}
