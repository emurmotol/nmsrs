package img

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"fmt"

	"errors"

	"io/ioutil"

	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

var (
	mimes            = []string{"image/jpeg", "image/png", "image/gif"}
	ErrImageNotValid = errors.New("We only support PNG, GIF, or JPG pictures")
	ErrImageToLarge  = fmt.Errorf("Please select a picture smaller than %s", str.BytesForHumans(env.DefaultMaxImageUploadSize))
)

func Save(file multipart.File, handler *multipart.FileHeader, name string) error {
	defer file.Close()
	dir := filepath.Dir(name)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(name, data, 0644)

	if err != nil {
		return err
	}
	return nil
} // TODO: Reuse code to separate package

func Validate(newFileInstance multipart.File, newHandlerInstance *multipart.FileHeader) error {
	for _, mime := range mimes {
		if newHandlerInstance.Header.Get("Content-Type") == mime {
			size, err := getSize(newFileInstance)

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
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return 0, err
	}
	return int64(len(data)), nil
} // TODO: Breaks the image, reuse code to separate package
