package civilstatus

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
		for _, value := range data {
			var civStat CivilStatus
			civStat.Name = value
			_, err := civStat.Insert()

			if err != nil {
				panic(err)
			}
		}
	}
}
