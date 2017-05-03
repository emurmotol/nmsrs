package eligibility

import "log"

var data = []string{
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

func Seeder() {
	eligs, err := All()

	if err != nil {
		panic(err)
	}

	if len(eligs) == 0 {
		for _, value := range data {
			var elig Eligibility
			elig.Name = value
			_, err := elig.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Eligibility seeding successful")
	}
}
