package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func TechnicalTrainingAndRelevantExperience(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Technical Training And Relevant Experience",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.technical_training_and_relevant_experience", data, funcMap)
}