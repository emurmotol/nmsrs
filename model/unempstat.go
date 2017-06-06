package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type UnEmpStat struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
}

func unEmpStatSeeder() {
	data := []string{
		"NEW ENTRANT/FRESH GRADUATE",
		"FINISHED CONTRACT",
		"RESIGNED",
		"TERMINATED/LAID OFF, LOCAL",
		"TERMINATED/LAID OFF, OVERSEAS",
	}

	for _, name := range data {
		unEmpStat := UnEmpStat{Name: strings.ToUpper(name)}
		unEmpStat.Create()
	}
}

func (unEmpStat *UnEmpStat) Create() *UnEmpStat {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&unEmpStat).Error; err != nil {
		panic(err)
	}
	return unEmpStat
}
