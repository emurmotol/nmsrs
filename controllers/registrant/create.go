package registrant

import (
	"net/http"
	"strconv"

	"github.com/zneyrl/nmsrs/helpers/res"
	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create registrant",
		// "EmploymentStatusOptions": models.EmploymentStatusOptions,
		// "SexOptions":              models.SexOptions,
		// "CivilStatusOptions":      models.CivilStatusOptions,
		// "ReligionOptions":         models.ReligionOptions,
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
