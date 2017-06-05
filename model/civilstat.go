package model

import "github.com/emurmotol/nmsrs/database"

type CivilStat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CivilStatSeeder() {
	data := []string{
		"SINGLE",
		"WIDOWED",
		"MARRIED",
		"SEPARATED",
		"OTHER",
	}

	for _, name := range data {
		civilStat := CivilStat{Name: name}

		if _, err := civilStat.Create(); err != nil {
			panic(err)
		}
	}
}

func (civilStat *CivilStat) Create() (*CivilStat, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&civilStat).Error; err != nil {
		return nil, err
	}
	return civilStat, nil
}
