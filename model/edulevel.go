package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type EduLevel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func eduLevelSeeder() {
	data := []string{
		"1ST YEAR COLLEGE LEVEL",
		"1ST YEAR HIGH SCHOOL/GRADE VII (FOR K TO 12)",
		"2ND YEAR COLLEGE LEVEL",
		"2ND YEAR HIGH SCHOOL/GRADE VIII (FOR K TO 12)",
		"3RD YEAR COLLEGE LEVEL",
		"3RD YEAR HIGH SCHOOL/GRADE IX (FOR K TO 12)",
		"4TH YEAR COLLEGE LEVEL",
		"4TH YEAR HIGH SCHOOL/GRADE X (FOR K TO 12)",
		"5TH YEAR COLLEGE LEVEL",
		"COLLEGE GRADUATE",
		"ELEMENTARY GRADUATE",
		"GRADE I",
		"GRADE II",
		"GRADE III",
		"GRADE IV",
		"GRADE V",
		"GRADE VI",
		"GRADE VII",
		"GRADE VIII",
		"GRADE XI (FOR K TO 12)",
		"GRADE XII (FOR K TO 12)",
		"HIGH SCHOOL GRADUATE",
		"MASTERAL/POST GRADUATE LEVEL",
		"MASTERAL/POST GRADUATE",
		"VOCATIONAL GRADUATE",
		"VOCATIONAL UNDERGRADUATE",
	}

	for _, name := range data {
		eduLevel := EduLevel{Name: strings.ToUpper(name)}
		eduLevel.Create()
	}
}

func (eduLevel *EduLevel) Create() *EduLevel {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&eduLevel).Error; err != nil {
		panic(err)
	}
	return eduLevel
}

func (eduLevel EduLevel) Index(q string) []EduLevel {
	db := database.Conn()
	defer db.Close()

	eduLevels := []EduLevel{}
	results := make(chan []EduLevel)

	go func() {
		db.Find(&eduLevels, "name LIKE ?", database.WrapLike(q))
		results <- eduLevels
	}()
	return <-results
}
