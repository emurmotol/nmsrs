package registrant

import (
	"net/http"
	"strconv"

	"github.com/zneyrl/nmsrs/helpers/res"
	"github.com/zneyrl/nmsrs/helpers/tpl"
	"github.com/zneyrl/nmsrs/models/civilstatus"
	"github.com/zneyrl/nmsrs/models/employmentstatus"
	"github.com/zneyrl/nmsrs/models/religion"
	"github.com/zneyrl/nmsrs/models/sex"
)

func Create(w http.ResponseWriter, r *http.Request) {
	empStats, err := employmentstatus.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	sexs, err := sex.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	civStats, err := civilstatus.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	religs, err := religion.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	data := map[string]interface{}{
		"Title":            "Create registrant",
		"EmploymentStatus": empStats,
		"Sex":              sexs,
		"CivilStatus":      civStats,
		"Religion":         religs,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "wizard", "registrant.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
	step, err := strconv.Atoi(r.URL.Query()["step"][0])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	switch step {
	case 1:
		res.JSON(w, res.Make{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"proceed": true,
			},
			Errors: "",
		})
		return
	case 2:
		return
	case 3:
		return
	case 4:
		return
	case 5:
		return
	case 6:
		return
	case 7:
		return
	case 8:
		return
	case 9:
		return
	default:
		return
	}
}
