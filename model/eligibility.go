package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Eligibility struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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
		eligibility := Eligibility{Name: strings.ToUpper(name)}

		if _, err := eligibility.Create(); err != nil {
			panic(err)
		}
	}
}

func (eligibility *Eligibility) Create() (*Eligibility, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&eligibility).Error; err != nil {
		return nil, err
	}
	return eligibility, nil
}

func (eligibility Eligibility) Search(q string) []Eligibility {
	db := database.Conn()
	defer db.Close()

	eligibilities := []Eligibility{}
	results := make(chan []Eligibility)

	go func() {
		db.Find(&eligibilities, "name LIKE ?", database.WrapLike(q))
		results <- eligibilities
	}()
	return <-results
}
