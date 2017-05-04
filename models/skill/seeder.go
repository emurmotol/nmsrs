package skill

import "log"

var data = []string{
	"Skill",
	"COMPUTER LITERATE",
	"DRIVER",
	"AUTO MECHANIC",
	"CARPENTRY WORK",
	"MASONRY",
	"ELECTRICIAN",
	"STENOGRAPHY",
	"PAINTING JOBS",
	"EMBROIDERY",
	"SEWING DRESSES",
	"TAILORING",
	"BEAUTICIAN",
	"DOMESTIC CHORES",
	"GARDENING",
	"PHOTOGRAPHY",
	"PAINTER/ARTIST",
}

func Seeder() {
	skills, err := All()

	if err != nil {
		panic(err)
	}

	if len(skills) == 0 {
		for _, value := range data {
			var skill Skill
			skill.Name = value
			_, err := skill.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Skill seeded")
	}
}
