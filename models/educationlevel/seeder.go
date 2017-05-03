package educationlevel

var data = []string{
	"GRADE I",
	"GRADE II",
	"GRADE III",
	"GRADE IV",
	"GRADE V",
	"GRADE VI",
	"GRADE VII",
	"GRADE VIII",
	"ELEMENTARY GRADUATE",
	"1ST YEAR HIGH SCHOOL/GRADE VII (FOR K TO 12)",
	"2ND YEAR HIGH SCHOOL/GRADE VIII (FOR K TO 12)",
	"3RD YEAR HIGH SCHOOL/GRADE IX (FOR K TO 12)",
	"4TH YEAR HIGH SCHOOL/GRADE X (FOR K TO 12)",
	"GRADE XI (FOR K TO 12)",
	"GRADE XII (FOR K TO 12)",
	"HIGH SCHOOL GRADUATE",
	"VOCATIONAL UNDERGRADUATE",
	"VOCATIONAL GRADUATE",
	"1ST YEAR COLLEGE LEVEL",
	"2ND YEAR COLLEGE LEVEL",
	"3RD YEAR COLLEGE LEVEL",
	"4TH YEAR COLLEGE LEVEL",
	"5TH YEAR COLLEGE LEVEL",
	"COLLEGE GRADUATE",
	"MASTERAL/POST GRADUATE LEVEL",
	"MASTERAL/POST GRADUATE",
}

func Seeder() {
	edulvls, err := All()

	if err != nil {
		panic(err)
	}

	if len(edulvls) == 0 {
		for _, value := range data {
			var edulvl EducationLevel
			edulvl.Name = value
			_, err := edulvl.Insert()

			if err != nil {
				panic(err)
			}
		}
	}
}
