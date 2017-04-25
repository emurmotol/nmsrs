package check

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs/helpers/img"
	"github.com/zneyrl/nmsrs/helpers/res"
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
	photoFieldName := mux.Vars(r)["field"]
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

	if photo != nil {
		defer photo.Close()

		if err := img.Validate(photo, handler); err != nil {
			if err == img.ErrImageNotValid || err == img.ErrImageToLarge { // TODO: Add new custom err here
				res.JSON(w, res.Make{
					Status: http.StatusForbidden,
					Data:   "",
					Errors: map[string]string{
						photoFieldName: err.Error(),
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
	}

	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": "Image is valid",
		},
		Errors: "",
	})
	return
}
