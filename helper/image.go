package helper

import (
	"errors"
	"fmt"
	"image"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/lang"
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
				return fmt.Errorf(lang.Get("imageTooLarge"), strconv.Itoa(maxSize))
			}
			return nil
		}
	}
	return errors.New(lang.Get("imageInvalid"))
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
