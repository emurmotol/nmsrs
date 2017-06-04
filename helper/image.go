package helper

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/emurmotol/nmsrs.v4/lang"
)

func ValidateImage(header *multipart.FileHeader) error {
	mimes, _ := env.Conf.List("default.photo.mimes")
	maxSize, _ := env.Conf.Int("default.photo.maxSize")

	for _, mime := range mimes {
		if strings.ToLower(header.Header.Get("Content-Type")) == mime.(string) {
			image, err := header.Open()

			if err != nil {
				return err
			}
			size, err := getFileSize(image)

			if err != nil {
				return err
			}

			if size > int64(maxSize*1048576) {
				return fmt.Errorf(lang.Get("image_too_large"), strconv.Itoa(maxSize))
			}
			return nil
		}
	}
	return errors.New(lang.Get("image_invalid"))
}

func SaveAsJPEG(file multipart.File, name string) error {
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

func crop(img image.Image) *image.NRGBA {
	var size int

	if img.Bounds().Dx() == img.Bounds().Dy() {
		size = img.Bounds().Dx()
	} else {
		size = minInt(img.Bounds().Dx(), img.Bounds().Dy())
	}
	dst := imaging.CropCenter(img, size, size)
	return imaging.Resize(dst, 256, 256, imaging.Lanczos)
}
