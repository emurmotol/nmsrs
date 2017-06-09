package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type OtherSkill struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
}

func otherSkillSeeder() {
	data := []string{
		"ANALYZING",
		"ANTICIPATING",
		"ASSEMBLING",
		"BUILDING",
		"CHECKING",
		"COMPARING",
		"COMPILING",
		"COMPUTING",
		"COORDINATING",
		"COPYING",
		"CREATING/INVENTING",
		"DISCOVERING",
		"DIVERTING",
		"DRIVING/STEERING",
		"ENCOURAGING",
		"EXPRESSING",
		"FEEDING/LOADING",
		"HELPING",
		"IMPLEMENTING",
		"INNOVATING",
		"INSPECTING",
		"INSTRUCTING",
		"INTERPRETING",
		"LEADING",
		"MACHINE WORK",
		"MANIPULATING",
		"MATERIALS HANDLING",
		"MOTIVATING",
		"NEGOTIATING",
		"OPERATING/CONTROLLING",
		"ORGANIZING",
		"PERSUADING",
		"PLANNING",
		"POSTING",
		"PRECISION WORKING",
		"PREDICTING",
		"PRODUCING",
		"PROMOTING",
		"RECORDING",
		"REPAIRING/ADJUSTING",
		"RESEARCHING",
		"SELLING",
		"SERVING",
		"SETTING - UP",
		"SETTING - UP/RESTORING",
		"SPEAKING",
		"SPECULATING",
		"SYNTHESIZING IDEA",
		"TABULATING",
		"TEACHING",
		"TESTING",
		"THEORIZING",
		"WAREHOUSING",
	}

	for _, name := range data {
		otherSkill := OtherSkill{Name: strings.ToUpper(name)}
		otherSkill.Create()
	}
}

func (otherSkill *OtherSkill) Create() *OtherSkill {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&otherSkill).Error; err != nil {
		panic(err)
	}
	return otherSkill
}

func (otherSkill OtherSkill) Index(q string) []OtherSkill {
	db := database.Conn()
	defer db.Close()

	otherSkills := []OtherSkill{}
	results := make(chan []OtherSkill)

	go func() {
		db.Find(&otherSkills, "name LIKE ?", database.WrapLike(q))
		results <- otherSkills
	}()

	otherSkills = <-results
	close(results)
	return otherSkills
}
