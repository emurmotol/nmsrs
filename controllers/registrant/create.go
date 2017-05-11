package registrant

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/helpers/tpl"
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
		panic(err)
	}
	sexs, err := sex.All()

	if err != nil {
		panic(err)
	}
	civStats, err := civilstatus.All()

	if err != nil {
		panic(err)
	}
	disabs, err := disability.All()

	if err != nil {
		panic(err)
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
		panic(err)
	}
	defer r.Body.Close()
	log.Printf("%v\n", t)

	step, err := strconv.Atoi(r.URL.Query().Get("step"))

	if err != nil {
		panic(err)
	}

	switch step {
	case 1:
		var personalInfo registrant.PersonalInformation

		if err := mapstructure.Decode(t["personal_information"], &personalInfo); err != nil {
			panic(err)
		}
		var basicInfo registrant.BasicInformation

		if err := mapstructure.Decode(t["basic_information"], &basicInfo); err != nil {
			panic(err)
		}
		res.JSON(w, res.Make{
			Status: http.StatusOK,
			Data: map[string]interface{}{
				"proceed": false,
			},
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
