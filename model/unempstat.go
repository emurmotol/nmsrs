package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type UnEmpStat struct {
	ID   int    `json:"id"`
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

		if _, err := unEmpStat.Create(); err != nil {
			panic(err)
		}
	}
}

func (unEmpStat *UnEmpStat) Create() (*UnEmpStat, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&unEmpStat).Error; err != nil {
		return nil, err
	}
	return unEmpStat, nil
}
