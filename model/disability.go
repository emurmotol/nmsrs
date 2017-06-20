package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Disability struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func disabilitySeeder() {
	data := []string{
		"VISUAL IMPAIRMENT",
		"HEARING IMPAIRMENT",
		"SPEECH IMPAIRMENT",
		"PHYSICAL IMPAIRMENT",
		"OTHER",
	}

	for _, name := range data {
		disability := Disability{Name: strings.ToUpper(name)}
		disability.Create()
	}
}

func (disability *Disability) Create() *Disability {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&disability).Error; err != nil {
		panic(err)
	}
	return disability
}

func Disabilities() []Disability {
	db := database.Con()
	defer db.Close()
	disabilities := []Disability{}

	if err := db.Find(&disabilities).Error; err != nil {
		panic(err)
	}
	return disabilities
}
