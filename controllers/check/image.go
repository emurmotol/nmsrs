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
	photoFieldName := mux.Vars(r)["id"]
	photo, handler, err := r.FormFile(photoFieldName)

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

	if err := img.Validate(photo, handler.Header); err != nil {
		if err == img.ErrImageNotValid || err == img.ErrImageToLarge { // TODO: Add new custom err here
			res.JSON(w, res.Make{
				Status: http.StatusInternalServerError,
				Data:   "",
				Errors: map[string]string{
					photoFieldName: fmt.Sprintf("%s %s", str.SnakeCaseToSentenceCase(photoFieldName), err.Error()),
				},
			})
			return
		}
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
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
