package fi

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	KB = int64(1024)
	MB = int64(1 << 20)
	GB = int64(KB * MB)
	TB = int64(KB * GB)
)

type Sizer interface {
	Size() int64
}

func Save(file multipart.File, name string) error {
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
}

func Size(file multipart.File) (int64, error) {
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return 0, err
	}
	return int64(len(data)), nil
} // TODO: Breaks the image, reuse code to separate package
