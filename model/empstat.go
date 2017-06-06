package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type EmpStat struct {
	ID   int    `json:"id"`
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

		if _, err := empStat.Create(); err != nil {
			panic(err)
		}
	}
}

func (empStat *EmpStat) Create() (*EmpStat, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&empStat).Error; err != nil {
		return nil, err
	}
	return empStat, nil
}
