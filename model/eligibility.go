package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Eligibility struct {
	ID   uint   `json:"id"`
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
		eligibility.Create()
	}
}

func (eligibility *Eligibility) Create() *Eligibility {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&eligibility).Error; err != nil {
		panic(err)
	}
	return eligibility
}

func (eligibility Eligibility) Index(q string) []Eligibility {
	db := database.Con()
	defer db.Close()

	eligibilities := []Eligibility{}
	results := make(chan []Eligibility)

	go func() {
		db.Find(&eligibilities, "name LIKE ?", database.WrapLike(q))
		results <- eligibilities
	}()
	eligibilities = <-results
	close(results)
	return eligibilities
}
