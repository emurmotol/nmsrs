package region

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type RefRegion struct {
	ID       int    `json:"id"`
	PsgcCode string `json:"psgcCode"`
	RegDesc  string `json:"regDesc"`
	RegCode  string `json:"regCode"`
}

func Seeder() {
	regs, err := All()

	if err != nil {
		panic(err)
	}

	if len(regs) == 0 {
		refRegs, err := ioutil.ReadFile("models/region/refregion.json")

		if err != nil {
			panic(err)
		}
		var rrs []RefRegion

		if err := json.Unmarshal(refRegs, &rrs); err != nil {
			panic(err)
		}

		for _, rr := range rrs {
			var reg Region
			reg.ID = rr.ID
			reg.Code = rr.RegCode
			reg.Desc = strings.ToUpper(rr.RegDesc)
			reg.PSGCCode = rr.PsgcCode
			_, err := reg.Insert()

			if err != nil {
				panic(err)
			}
		}
		log.Println("Region seeded")
	}
}
