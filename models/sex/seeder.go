package sex

import "log"

var data = []string{
	"MALE",
	"FEMALE",
}

func Seeder() {
	sexs, err := All()

	if err != nil {
		panic(err)
	}

	if len(sexs) == 0 {
		for index, value := range data {
			var sex Sex
			sex.ID = index + 1
			sex.Name = value
			_, err := sex.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Sex seeded")
	}
}
