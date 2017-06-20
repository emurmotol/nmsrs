package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type CivilStat struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
}

func civilStatSeeder() {
	data := []string{
		"SINGLE",
		"WIDOWED",
		"MARRIED",
		"SEPARATED",
		"OTHER",
	}

	for _, name := range data {
		civilStat := CivilStat{Name: strings.ToUpper(name)}
		civilStat.Create()
	}
}

func (civilStat *CivilStat) Create() *CivilStat {
	db := database.Con()
	defer db.Close()

	if err := db.Create(&civilStat).Error; err != nil {
		panic(err)
	}
	return civilStat
}

func CivilStats() []CivilStat {
	db := database.Con()
	defer db.Close()
	civilStats := []CivilStat{}

	if err := db.Find(&civilStats).Error; err != nil {
		panic(err)
	}
	return civilStats
}
