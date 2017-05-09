package otherskill

import (
	"log"
	"strings"
)

var data = []string{
	"TEACHING",
	"NEGOTIATING",
	"DIVERTING",
	"PERSUADING",
	"SPEAKING",
	"SERVING",
	"HELPING",
	"ENCOURAGING",
	"MOTIVATING",
	"LEADING",
	"PROMOTING",
	"SELLING",
	"COORDINATING",
	"ANALYZING",
	"COMPILING",
	"COMPUTING",
	"TABULATING",
	"COMPARING",
	"PLANNING",
	"RECORDING",
	"POSTING",
	"CHECKING",
	"RESEARCHING",
	"TESTING",
	"COPYING",
	"MACHINE WORK",
	"SETTING - UP",
	"OPERATING/CONTROLLING",
	"DRIVING/STEERING",
	"MANIPULATING",
	"MATERIALS HANDLING",
	"INSPECTING",
	"PRODUCING",
	"WAREHOUSING",
	"BUILDING",
	"PRECISION WORKING",
	"SETTING - UP/RESTORING",
	"FEEDING/LOADING",
	"ASSEMBLING",
	"REPAIRING/ADJUSTING",
	"IMPLEMENTING",
	"SYNTHESIZING IDEA",
	"CREATING/INVENTING",
	"DISCOVERING",
	"INTERPRETING",
	"EXPRESSING",
	"INSTRUCTING",
	"ORGANIZING",
	"THEORIZING",
	"SPECULATING",
	"PREDICTING",
	"ANTICIPATING",
	"INNOVATING",
}

func Seeder() {
	oskills, err := All()

	if err != nil {
		panic(err)
	}

	if len(oskills) == 0 {
		for index, value := range data {
			var oskill OtherSkill
			oskill.ID = index + 1
			oskill.Name = strings.ToUpper(value)
			_, err := oskill.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("OtherSkill seeded")
	}
}
