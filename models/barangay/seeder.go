package barangay

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type RefBrgy struct {
	BrgyCode    string `json:"brgyCode"`
	BrgyDesc    string `json:"brgyDesc"`
	RegCode     string `json:"regCode"`
	ProvCode    string `json:"provCode"`
	CityMunCode string `json:"cityMunCode"`
}

func Seeder() {
	brgys, err := All()

	if err != nil {
		panic(err)
	}

	if len(brgys) == 0 {
		refBrgys, err := ioutil.ReadFile("models/barangay/refbrgy.json")

		if err != nil {
			panic(err)
		}
		var rbs []RefBrgy

		if err := json.Unmarshal(refBrgys, &rbs); err != nil {
			panic(err)
		}

		for _, rb := range rbs {
			var brgy Barangay
			brgy.Code = rb.BrgyCode
			brgy.Desc = strings.ToUpper(rb.BrgyDesc)
			brgy.RegionCode = rb.RegCode
			brgy.ProvinceCode = rb.ProvCode
			brgy.CityMunicipalityCode = rb.CityMunCode
			_, err := brgy.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Barangay seeded")
	}
}
