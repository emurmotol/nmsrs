package helper

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

func getFileSize(file multipart.File) (int64, error) {
	f := file.(*os.File)
	fi, err := f.Stat()

	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func SaveAsJson(data []interface{}, name string) error {
	j, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fmt.Sprintf("import/%s.json", name), j, 0644); err != nil {
		return err
	}
	return nil
}

func SaveAsJpeg(file multipart.File, name string) error {
	defer file.Close()
	src, _, err := image.Decode(file)

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

	if err := jpeg.Encode(f, crop(src), &opt); err != nil {
		return err
	}
	return nil
}
