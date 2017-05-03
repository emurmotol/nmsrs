package employmentstatus

var data = []string{
	"WAGED EMPLOYED",
	"SELF EMPLOYED",
	"UNEMPLOYED",
}

func Seeder() {
	empstats, err := All()

	if err != nil {
		panic(err)
	}

	if len(empstats) == 0 {
		for _, value := range data {
			var empstat EmploymentStatus
			empstat.Name = value
			_, err := empstat.Insert()

			if err != nil {
				panic(err)
			}
		}
	}
}
