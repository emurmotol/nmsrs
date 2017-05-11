package check

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/img"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/gorilla/mux"
)

func Image(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	photoFieldName := mux.Vars(r)["field"]
	photo, handler, err := r.FormFile(photoFieldName)

	if err != http.ErrMissingFile {
		panic(err)
	}

	if photo != nil {
		defer photo.Close()

		if err := img.Validate(photo, handler); err != nil {
			if err == img.ErrImageNotValid || err == img.ErrImageTooLarge { // TODO: Add new custom err here
				res.JSON(w, res.Make{
					Status: http.StatusForbidden,
					Data: map[string]string{
						"error": err.Error(),
					},
				})
				return
			}
			panic(err)
		}
	}

	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": lang.En["image_valid"],
		},
	})
	return
}
