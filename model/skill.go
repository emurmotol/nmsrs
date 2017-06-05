package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SkillSeeder() {
	data := []string{
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

	for _, name := range data {
		skill := Skill{Name: strings.ToUpper(name)}

		if _, err := skill.Create(); err != nil {
			panic(err)
		}
	}
}

func (skill *Skill) Create() (*Skill, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (skill Skill) Search(q string) []Skill {
	db := database.Conn()
	defer db.Close()

	skills := []Skill{}
	results := make(chan []Skill)

	go func() {
		db.Find(&skills, "name LIKE ?", database.WrapLike(q))
		results <- skills
	}()
	return <-results
}
