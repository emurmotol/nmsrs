package db

import (
	"fmt"
	"log"

	"github.com/zneyrl/nmsrs/env"
	mgo "gopkg.in/mgo.v2"
)

var (
	DB                 *mgo.Database
	Certificates       *mgo.Collection
	CivilStatuses      *mgo.Collection
	Countries          *mgo.Collection
	Courses            *mgo.Collection
	Disabilities       *mgo.Collection
	EducationLevels    *mgo.Collection
	Eligibilities      *mgo.Collection
	EmploymentStatuses *mgo.Collection
	Industries         *mgo.Collection
	Languages          *mgo.Collection
	Licenses           *mgo.Collection
	OtherSkills        *mgo.Collection
	Positions          *mgo.Collection
	Registrants        *mgo.Collection
	Religions          *mgo.Collection
	Schools            *mgo.Collection
	Sexes              *mgo.Collection
	Skills             *mgo.Collection
	UnemployedStatuses *mgo.Collection
	Users              *mgo.Collection
)

func init() {
	s, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%d", env.DBHost, env.DBPort))

	if err != nil {
		panic(err)
	}

	if err := s.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB connection successful")

	DB = s.DB(env.DBName)
	Certificates = DB.C("certificates")
	CivilStatuses = DB.C("civilStatuses")
	Countries = DB.C("countries")
	Courses = DB.C("courses")
	Disabilities = DB.C("disabilities")
	EducationLevels = DB.C("educationLevels")
	Eligibilities = DB.C("eligibilities")
	EmploymentStatuses = DB.C("employmentStatuses")
	Industries = DB.C("industries")
	Languages = DB.C("languages")
	Licenses = DB.C("licenses")
	OtherSkills = DB.C("otherSkills")
	Positions = DB.C("positions")
	Registrants = DB.C("registrants")
	Religions = DB.C("religions")
	Schools = DB.C("schools")
	Sexes = DB.C("sexes")
	Skills = DB.C("skills")
	UnemployedStatuses = DB.C("unemployedStatuses")
	Users = DB.C("users")
}
