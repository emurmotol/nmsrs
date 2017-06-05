package model

import "github.com/emurmotol/nmsrs/database"

type Industry struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func IndustrySeeder() {
	data := []string{
		"AGRICULTURE",
		"FISHING",
		"MINING AND QUARRYING",
		"MANUFACTURING",
		"ELECTRICITY, GAS AND WATER SUPPLY",
		"CONSTRUCTION",
		"WHOLESALE AND RETAIL TRADE",
		"HOTELS AND RESTAURANTS",
		"TRANSPORT, STORAGE AND COMMUNICATION",
		"FINANCIAL INTERMEDIATION",
		"REAL ESTATE, RENTING AND BUSINESS ACTIVITIES",
		"PUBLIC ADMINISTRATION AND DEFENSE",
		"EDUCATION",
		"HEALTH AND SOCIAL WORK",
		"OTHER COMMUNITY, SOCIAL AND PERSONAL SERVICE ACTIVITIES",
		"ACTIVITIES OF PRIVATE HOUSEHOLDS AS EMPLOYERS AND UNDIFFENTIATED PRODUCTION ACTIVITIES OF PRIVATE",
		"EXTRA - TERRITORIAL ORGANIZATIONS AND BODIES",
	}

	for _, name := range data {
		industry := Industry{Name: name}

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
