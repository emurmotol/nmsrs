package barangay

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/barangay"
	"github.com/gorilla/mux"
)

func All(w http.ResponseWriter, r *http.Request) {
	brgys, err := barangay.Search(bson.M{
		"cityMunicipalityCode": mux.Vars(r)["city_municipality_code"],
		"desc": bson.RegEx{
			Pattern: r.URL.Query().Get("q"),
			Options: "i",
		},
	})

	if err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"barangays": brgys,
		},
	})
}
