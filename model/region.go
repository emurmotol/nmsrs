package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Region struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Desc     string `json:"desc"`
	PsgcCode string `json:"psgc_code"`
}

type RefRegion struct {
	PsgcCode string `json:"psgcCode"`
	RegDesc  string `json:"regDesc"`
	RegCode  string `json:"regCode"`
}

func RegionSeeder() {
	data, err := ioutil.ReadFile("model/data/refregion.json")

	if err != nil {
		panic(err)
	}
	refRegions := []RefRegion{}

	if err := json.Unmarshal(data, &refRegions); err != nil {
		panic(err)
	}

	for _, refRegion := range refRegions {
		region := Region{
			Code:     refRegion.RegCode,
			Desc:     strings.ToUpper(refRegion.RegDesc),
			PsgcCode: refRegion.PsgcCode,
		}

		if _, err := region.Create(); err != nil {
			panic(err)
		}
	}
}

func (region *Region) Create() (*Region, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&region).Error; err != nil {
		return nil, err
	}
	return region, nil
}
