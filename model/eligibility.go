package model

import "github.com/emurmotol/nmsrs/database"

type Eligibility struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func EligibilitySeeder() {
	data := []string{
		"CAREER SERVICE PROFESSIONAL",
		"CAREER SERVICE EXECUTIVE ELIGIBILITY",
		"CAREER EXECUTIVE OFFICER ELIGIBILITY",
		"R.A. 1080",
		"CAREER SERVICE SUB - PROFESSIONAL",
		"CAREER EXECUTIVE SERVICE OFFICER",
		"POLICE OFFICER 1",
		"STENOGRAPHER",
		"SOIL TECHNOLOGIST",
		"DATA ENCODER",
		"FORESTRY EXTENSION SERVICE",
		"FIRE OFFICER 2",
	}

	for _, name := range data {
		eligibility := Eligibility{Name: name}

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
