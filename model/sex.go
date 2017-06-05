package model

import "github.com/emurmotol/nmsrs/database"

type Sex struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SexSeeder() {
	data := []string{
		"MALE",
	"FEMALE",
	}

	for _, name := range data {
		sex := Sex{Name: name}

		if _, err := sex.Create(); err != nil {
			panic(err)
		}
	}
}

func (sex *Sex) Create() (*Sex, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&sex).Error; err != nil {
		return nil, err
	}
	return sex, nil
}
