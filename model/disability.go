package model

import "github.com/emurmotol/nmsrs/database"

type Disability struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func DisabilitySeeder() {
	data := []string{
		"VISUAL IMPAIRMENT",
		"HEARING IMPAIRMENT",
		"SPEECH IMPAIRMENT",
		"PHYSICAL IMPAIRMENT",
		"OTHER",
	}

	for _, name := range data {
		disability := Disability{Name: name}

		if _, err := disability.Create(); err != nil {
			panic(err)
		}
	}
}

func (disability *Disability) Create() (*Disability, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&disability).Error; err != nil {
		return nil, err
	}
	return disability, nil
}
