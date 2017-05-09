package api

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/country"
)

func Countries(w http.ResponseWriter, r *http.Request) {
	couns, err := country.Search(bson.M{
		"$and": []bson.M{
			bson.M{
				"name": bson.RegEx{
					Pattern: r.URL.Query().Get("q"),
					Options: "i",
				},
			},
			bson.M{
				"name": bson.M{
					"$ne": r.URL.Query().Get("except"),
				},
			},
		},
	})

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
			"countries": couns,
		},
		Errors: "",
	})
}
