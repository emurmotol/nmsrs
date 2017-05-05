package api

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/barangay"
	"github.com/gorilla/mux"
)

func Barangays(w http.ResponseWriter, r *http.Request) {
	brgys, err := barangay.FindAllBy("cityMunicipalityCode", mux.Vars(r)["city_municipality_code"])

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
			"barangays": brgys,
		},
		Errors: "",
	})
}