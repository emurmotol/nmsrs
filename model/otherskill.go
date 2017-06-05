package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type OtherSkill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func OtherSkillSeeder() {
	data := []string{
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

	for _, name := range data {
		otherSkill := OtherSkill{Name: strings.ToUpper(name)}

		if _, err := otherSkill.Create(); err != nil {
			panic(err)
		}
	}
}

func (otherSkill *OtherSkill) Create() (*OtherSkill, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&otherSkill).Error; err != nil {
		return nil, err
	}
	return otherSkill, nil
}

func (otherSkill OtherSkill) Search(q string) []OtherSkill {
	db := database.Conn()
	defer db.Close()

	otherSkills := []OtherSkill{}
	results := make(chan []OtherSkill)

	go func() {
		db.Find(&otherSkills, "name LIKE ?", database.WrapLike(q))
		results <- otherSkills
	}()
	return <-results
}
