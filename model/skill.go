package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func skillSeeder() {
	data := []string{
		"AUTO MECHANIC",
		"BEAUTICIAN",
		"CARPENTRY WORK",
		"COMPUTER LITERATE",
		"DOMESTIC CHORES",
		"DRIVER",
		"ELECTRICIAN",
		"EMBROIDERY",
		"GARDENING",
		"MASONRY",
		"PAINTER/ARTIST",
		"PAINTING JOBS",
		"PHOTOGRAPHY",
		"SEWING DRESSES",
		"STENOGRAPHY",
		"TAILORING",
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
