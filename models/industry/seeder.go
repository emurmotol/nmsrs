package industry

import (
	"log"
	"strings"
)

var data = []string{
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

func Seeder() {
	inds, err := All()

	if err != nil {
		panic(err)
	}

	if len(inds) == 0 {
		for index, value := range data {
			var ind Industry
			ind.ID = index + 1
			ind.Name = strings.ToUpper(value)
			_, err := ind.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Industry seeded")
	}
}
