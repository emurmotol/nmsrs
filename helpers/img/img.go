package img

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zneyrl/nmsrs-lookup/env"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("not a valid image")
	ErrImageToLarge  = errors.New("to large") // TODO: Include size in mb
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
