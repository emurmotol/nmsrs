package img

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

// func mimeFromIncipit(incipit []byte) string {
// 	for magic, mime := range magicTable {
// 		if strings.HasPrefix(incipit, magic) {
// 			return mime
// 		}
// 	}
// 	return ""
// }

func ValidateAndSave(photo multipart.File, handler *multipart.FileHeader, id string) error {
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
