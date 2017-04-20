package img

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SetUserPhoto(photo multipart.File, handler *multipart.FileHeader, id string) error {
	defer photo.Close()
	name := fmt.Sprintf("default%s", strings.ToLower(filepath.Ext(handler.Filename)))
	path := fmt.Sprintf("content/%s/img", id)
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(path, 0777) // TODO: Change permission
	}
	filepath := filepath.Join(path, name)

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
