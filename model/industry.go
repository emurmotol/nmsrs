package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Industry struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func industrySeeder() {
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
		industry.Create()
	}
}

func (industry *Industry) Create() *Industry {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&industry).Error; err != nil {
		panic(err)
	}
	return industry
}

func (industry Industry) Index(q string) []Industry {
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
