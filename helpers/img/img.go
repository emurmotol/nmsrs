package img

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"fmt"

	"errors"

	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("We only support PNG, GIF, or JPG pictures")
	ErrImageToLarge  = fmt.Errorf("Please upload a picture smaller than %s", str.BytesForHumans(env.DefaultMaxImageUploadSize))
)

func Save(photo multipart.File, handler *multipart.FileHeader, path string) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, photo)

	if err != nil {
		return err
	}
	return nil
}

func Validate(photo multipart.File, handler *multipart.FileHeader) error {
	for _, mime := range mimes {
		if handler.Header.Get("Content-Type") == mime {
			size, err := getSize(photo)

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

func getSize(file multipart.File) (int64, error) {
	var buff bytes.Buffer
	size, err := buff.ReadFrom(file)

	if err != nil {
		return 0, err
	}
	return size, nil
}
