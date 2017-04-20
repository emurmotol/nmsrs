package img

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("not a valid image")
)

func Save(photo multipart.File, handler *multipart.FileHeader, dir string, filename string) error {
	defer photo.Close()
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777) // TODO: Change permission
	}
	filepath := filepath.Join(dir, filename)

	file, err := os.Create(filepath)
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

func Validate(contentType string) error {
	for _, mime := range mimes {
		if contentType == mime {
			return nil
		}
	}
	return ErrImageNotValid
}
