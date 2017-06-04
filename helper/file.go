package helper

import (
	"mime/multipart"
	"os"
)

func getFileSize(file multipart.File) (int64, error) {
	f := file.(*os.File)
	fi, err := f.Stat()

	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
