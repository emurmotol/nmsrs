package disability

import "log"

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
		for index, value := range data {
			var disab Disability
			disab.ID = index + 1
			disab.Name = value
			_, err := disab.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Disability seeded")
	}
}
