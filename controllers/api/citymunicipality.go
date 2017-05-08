package api

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/citymunicipality"
)

func CityMunicipalitiesWithProvinces(w http.ResponseWriter, r *http.Request) {
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
	cityMunsWithProvs, err := citymunicipality.WithProvince(query)

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"city_municipalities_with_provinces": cityMunsWithProvs,
		},
		Errors: "",
	})
}
