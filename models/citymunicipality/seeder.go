package citymunicipality

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type RefCityMun struct {
	PsgcCode    string `json:"psgcCode"`
	CityMunDesc string `json:"cityMunDesc"`
	RegCode     string `json:"regCode"`
	ProvCode    string `json:"provCode"`
	CityMunCode string `json:"cityMunCode"`
}

func Seeder() {
	cityMuns, err := All()

	if err != nil {
		panic(err)
	}

	if len(cityMuns) == 0 {
		refCityMuns, err := ioutil.ReadFile("models/citymunicipality/refcitymun.json")

		if err != nil {
			panic(err)
		}
		var rcms []RefCityMun

		if err := json.Unmarshal(refCityMuns, &rcms); err != nil {
			panic(err)
		}

		for _, rcm := range rcms {
			var cityMun CityMunicipality
			cityMun.Code = rcm.CityMunCode
			cityMun.Desc = strings.ToUpper(rcm.CityMunDesc)
			cityMun.PSGCCode = rcm.PsgcCode
			cityMun.RegionCode = rcm.RegCode
			cityMun.ProvinceCode = rcm.ProvCode
			_, err := cityMun.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("CityMunicipality seeded")
	}
}
