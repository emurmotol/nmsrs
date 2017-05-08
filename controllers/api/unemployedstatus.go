package api

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/unemployedstatus"
)

func UnemployedStatuses(w http.ResponseWriter, r *http.Request) {
	unEmpStats, err := unemployedstatus.All()

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
			"unemployed_statuses": unEmpStats,
		},
		Errors: "",
	})
}
