package api

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/citymunicipality"
)

func CityMunicipalities(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	query := bson.M{
		"$or": []bson.M{
			bson.M{
				"desc": bson.RegEx{
					Pattern: q,
					Options: "i",
				},
			},
			bson.M{
				"province.desc": bson.RegEx{
					Pattern: q,
					Options: "i",
				},
			},
		},
	}
	cityMuns, err := citymunicipality.Search(query)

	if err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"city_municipalities": cityMuns,
		},
	})
}
