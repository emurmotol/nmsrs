package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Sex struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func sexSeeder() {
	data := []string{
		"MALE",
		"FEMALE",
	}

	for _, name := range data {
		sex := Sex{Name: strings.ToUpper(name)}
		sex.Create()
	}
}

func (sex *Sex) Create() *Sex {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&sex).Error; err != nil {
		panic(err)
	}
	return sex
}

func Sexes() []Sex {
	db := database.Con()
	defer db.Close()
	sexes := []Sex{}

	if err := db.Find(&sexes).Error; err != nil {
		panic(err)
	}
	return sexes
}
