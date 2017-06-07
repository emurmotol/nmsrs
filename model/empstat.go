package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type EmpStat struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func empStatSeeder() {
	data := []string{
		"WAGED EMPLOYED",
		"SELF EMPLOYED",
		"UNEMPLOYED",
	}

	for _, name := range data {
		empStat := EmpStat{Name: strings.ToUpper(name)}
		empStat.Create()
	}
}

func (empStat *EmpStat) Create() *EmpStat {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&empStat).Error; err != nil {
		panic(err)
	}
	return empStat
}

func EmpStats() []EmpStat {
	db := database.Conn()
	defer db.Close()
	empStats := []EmpStat{}

	if err := db.Find(&empStats).Error; err != nil {
		panic(err)
	}
	return empStats
}
