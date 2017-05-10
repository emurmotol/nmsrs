package registrant

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/helpers/vald"
	"github.com/emurmotol/nmsrs/models/civilstatus"
	"github.com/emurmotol/nmsrs/models/disability"
	"github.com/emurmotol/nmsrs/models/employmentstatus"
	"github.com/emurmotol/nmsrs/models/registrant"
	"github.com/emurmotol/nmsrs/models/sex"
	"github.com/mitchellh/mapstructure"
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
	disabs, err := disability.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	data := map[string]interface{}{
		"Title":              "Create registrant",
		"EmploymentStatuses": empStats,
		"Sexes":              sexs,
		"CivilStatuses":      civStats,
		"Disabilities":       disabs,
	}
	funcMap := map[string]interface{}{
		"SentenceCaseToSnakeCase": str.SentenceCaseToSnakeCase,
		"AllCapsToSentenceCase":   str.AllCapsToSentenceCase,
	}
	tpl.Render(w, r, "wizard", "registrant.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var t map[string]interface{}

	if err := d.Decode(&t); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	defer r.Body.Close()
	log.Printf("%v\n", t)

	step, err := strconv.Atoi(r.URL.Query().Get("step"))

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
		var personalInfo registrant.PersonalInformation

		if err := mapstructure.Decode(t["personal_information"], &personalInfo); err != nil {
			res.JSON(w, res.Make{
				Status: http.StatusInternalServerError,
				Data:   "",
				Errors: err.Error(),
			})
			return
		}
		errs := vald.StructHasError(personalInfo)
		var basicInfo registrant.BasicInformation

		if err := mapstructure.Decode(t["basic_information"], &basicInfo); err != nil {
			res.JSON(w, res.Make{
				Status: http.StatusInternalServerError,
				Data:   "",
				Errors: err.Error(),
			})
			return
		}

		for k, v := range vald.StructHasError(basicInfo) {
			if _, ok := errs[k]; !ok {
				errs[k] = v
			}
		}

		if len(basicInfo.CivilStatus) == 0 {
			if _, ok := errs["civil_status"]; !ok {
				errs["civil_status"] = "Civil status is a required field"
			}
		}

		if len(errs) != 0 {
			res.JSON(w, res.Make{
				Status: http.StatusForbidden,
				Data:   "",
				Errors: errs,
			})
			return
		}

		res.JSON(w, res.Make{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"proceed": false,
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
