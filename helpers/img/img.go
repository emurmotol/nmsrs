package img

import (
	"mime/multipart"

	"fmt"

	"errors"

	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/fi"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("We only support PNG, GIF, or JPG pictures")
	ErrImageToLarge  = fmt.Errorf("Please select a picture smaller than %s", str.BytesForHumans(env.DefaultMaxImageUploadSize))
)

func Validate(newFileInstance multipart.File, newHandlerInstance *multipart.FileHeader) error {
	for _, mime := range mimes {
		if newHandlerInstance.Header.Get("Content-Type") == mime {
			size, err := fi.Size(newFileInstance)

			if err != nil {
				return err
			}

			if size > env.DefaultMaxImageUploadSize {
				return ErrImageToLarge
			}
			return nil
		}
	}
	return ErrImageNotValid
}
