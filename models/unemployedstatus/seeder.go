package unemployedstatus

import "log"

var data = []string{
	"NEW ENTRANT/FRESH GRADUATE",
	"FINISHED CONTRACT",
	"RESIGNED",
	"TERMINATED/LAID OFF, LOCAL",
	"TERMINATED/LAID OFF, OVERSEAS",
}

func Seeder() {
	unEmpStats, err := All()

	if err != nil {
		panic(err)
	}

	if len(unEmpStats) == 0 {
		for index, value := range data {
			var unEmpStat UnemployedStatus
			unEmpStat.ID = index + 1
			unEmpStat.Name = value
			_, err := unEmpStat.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("UnemployedStatus seeded")
	}
}
