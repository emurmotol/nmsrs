package province

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type RefProvince struct {
	PsgcCode string `json:"psgcCode"`
	ProvDesc string `json:"provDesc"`
	RegCode  string `json:"regCode"`
	ProvCode string `json:"provCode"`
}

func Seeder() {
	provs, err := All()

	if err != nil {
		panic(err)
	}

	if len(provs) == 0 {
		refProvs, err := ioutil.ReadFile("models/province/refprovince.json")

		if err != nil {
			panic(err)
		}
		var rps []RefProvince

		if err := json.Unmarshal(refProvs, &rps); err != nil {
			panic(err)
		}

		for _, rp := range rps {
			var prov Province
			prov.Code = rp.ProvCode
			prov.Desc = strings.ToUpper(rp.ProvDesc)
			prov.PSGCCode = rp.PsgcCode
			prov.RegionCode = rp.RegCode
			_, err := prov.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Province seeded")
	}
}
