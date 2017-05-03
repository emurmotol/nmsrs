package unemployedstatus

var data = []string{
	"ACTIVELY LOOKING FOR WORK",
	"RESIGNED",
	"TERMINATED/LAID OFF, LOCAL",
	"TERMINATED/LAID OFF, OVERSEAS",
}

func Seeder() {
	unempstats, err := All()

	if err != nil {
		panic(err)
	}

	if len(unempstats) == 0 {
		for _, value := range data {
			var unempstat UnemployedStatus
			unempstat.Name = value
			_, err := unempstat.Insert()

			if err != nil {
				panic(err)
			}
		}
	}
}
