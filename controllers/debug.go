package controllers

import (
	"fmt"
	"net/http"

	"github.com/zneyrl/nmsrs/models/user"
)

func DD(w http.ResponseWriter, r *http.Request) {
	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddPage()
	// pdf.SetFont("Arial", "B", 16)
	// pdf.Cell(40, 10, "OK")
	// err := pdf.OutputFileAndClose("temp/hello.pdf")

	// if err != nil {
	// 	panic(err)
	// }
	// http.ServeFile(w, r, "temp/hello.pdf")
	val, _ := user.All()
	w.Write([]byte(fmt.Sprintf("%v\n", val)))
}
