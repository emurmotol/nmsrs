package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Industry struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndustrySeeder() {
	data := []string{
		"ACTIVITIES OF PRIVATE HOUSEHOLDS AS EMPLOYERS AND UNDIFFENTIATED PRODUCTION ACTIVITIES OF PRIVATE",
		"AGRICULTURE",
		"CONSTRUCTION",
		"EDUCATION",
		"ELECTRICITY, GAS AND WATER SUPPLY",
		"EXTRA - TERRITORIAL ORGANIZATIONS AND BODIES",
		"FINANCIAL INTERMEDIATION",
		"FISHING",
		"HEALTH AND SOCIAL WORK",
		"HOTELS AND RESTAURANTS",
		"MANUFACTURING",
		"MINING AND QUARRYING",
		"OTHER COMMUNITY, SOCIAL AND PERSONAL SERVICE ACTIVITIES",
		"PUBLIC ADMINISTRATION AND DEFENSE",
		"REAL ESTATE, RENTING AND BUSINESS ACTIVITIES",
		"TRANSPORT, STORAGE AND COMMUNICATION",
		"WHOLESALE AND RETAIL TRADE",
	}

	for _, name := range data {
		industry := Industry{Name: strings.ToUpper(name)}

		if _, err := industry.Create(); err != nil {
			panic(err)
		}
	}
}

func (industry *Industry) Create() (*Industry, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&industry).Error; err != nil {
		return nil, err
	}
	return industry, nil
}

func (industry Industry) Search(q string) []Industry {
	db := database.Conn()
	defer db.Close()

	industries := []Industry{}
	results := make(chan []Industry)

	go func() {
		db.Find(&industries, "name LIKE ?", database.WrapLike(q))
		results <- industries
	}()
	return <-results
}
