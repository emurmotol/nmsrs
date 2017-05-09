package civilstatus

import "log"

var data = []string{
	"SINGLE",
	"WIDOWED",
	"MARRIED",
	"SEPARATED",
	"OTHER",
}

func Seeder() {
	civStats, err := All()

	if err != nil {
		panic(err)
	}

	if len(civStats) == 0 {
		for index, value := range data {
			var civStat CivilStatus
			civStat.ID = index + 1
			civStat.Name = value
			_, err := civStat.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("CivilStatus seeded")
	}
}
