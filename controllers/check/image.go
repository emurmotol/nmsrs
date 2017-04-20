package check

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/helpers/img"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

func Image(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	_, handler, err := r.FormFile("photo")

	if err != nil {
		if err != http.ErrMissingFile {
			res.JSON(w, res.Make{
				Status: http.StatusInternalServerError,
				Data:   "",
				Errors: err.Error(),
			})
			return
		}
	}

	if err := img.Validate(handler.Header.Get("Content-Type")); err != nil {
		id := mux.Vars(r)["id"]
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: map[string]string{
				id: fmt.Sprintf("%s is %s", str.SnakeCaseToSentenceCase(id), err.Error()),
			},
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data:   "",
		Errors: "",
	})
	return
}
