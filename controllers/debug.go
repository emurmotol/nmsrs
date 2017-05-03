package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
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
	file, err := os.Open("temp/col.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	txt := ""

	for scanner.Scan() {
		txt += fmt.Sprintln(fmt.Sprintf(`"%s",`, scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%v\n", txt)))
}
