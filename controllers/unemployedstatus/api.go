package unemployedstatus

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/models/unemployedstatus"
)

func All(w http.ResponseWriter, r *http.Request) {
	unEmpStats, err := unemployedstatus.All()

	if err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"unemployed_statuses": unEmpStats,
		},
	})
}
