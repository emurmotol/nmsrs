package employmentstatus

import "log"

var data = []string{
	"WAGED EMPLOYED",
	"SELF EMPLOYED",
	"UNEMPLOYED",
}

func Seeder() {
	empStats, err := All()

	if err != nil {
		panic(err)
	}

	if len(empStats) == 0 {
		for _, value := range data {
			var empStat EmploymentStatus
			empStat.Name = value
			_, err := empStat.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("EmploymentStatus seeded")
	}
}
