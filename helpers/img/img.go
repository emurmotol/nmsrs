package img

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path/filepath"

	"fmt"

	"errors"

	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/helpers/fi"
	"github.com/zneyrl/nmsrs/helpers/str"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("We only support PNG, GIF, or JPG pictures")
	ErrImageToLarge  = fmt.Errorf("Please select a picture smaller than %s", str.BytesForHumans(env.DefaultMaxImageUploadSize))
)

func Validate(newFileInstance multipart.File, handler *multipart.FileHeader) error {
	for _, mime := range mimes {
		if handler.Header.Get("Content-Type") == mime {
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

func SaveAsJPEG(file multipart.File, name string) error {
	defer file.Close()
	photo, _, err := image.Decode(file)

	if err != nil {
		return err
	}
	dir := filepath.Dir(name)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	f, err := os.Create(name)

	if err != nil {
		return err
	}
	defer f.Close()
	var opt jpeg.Options
	opt.Quality = jpeg.DefaultQuality

	if err := jpeg.Encode(f, photo, &opt); err != nil {
		return err
	}
	return nil
}
