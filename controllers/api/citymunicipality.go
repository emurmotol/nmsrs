package api

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/citymunicipality"
)

func CityMunicipalitiesWithProvinces(w http.ResponseWriter, r *http.Request) {
	cityMunsWithProvs, err := citymunicipality.WithProvince()

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
