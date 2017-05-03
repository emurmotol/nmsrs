package disability

var data = []string{
	"VISUAL IMPAIRMENT",
	"HEARING IMPAIRMENT",
	"SPEECH IMPAIRMENT",
	"PHYSICAL IMPAIRMENT",
	"OTHER",
}

func Seeder() {
	disabs, err := All()

	if err != nil {
		panic(err)
	}

	if len(disabs) == 0 {
		for _, value := range data {
			var disab Disability
			disab.Name = value
			_, err := disab.Insert()

			if err != nil {
				panic(err)
			}
		}
	}
}
