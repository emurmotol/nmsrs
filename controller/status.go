package controller

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

func Forbidden(w http.ResponseWriter, r *http.Request) {
	renderStatus(w, http.StatusForbidden)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	renderStatus(w, http.StatusNotFound)
}

func renderStatus(w http.ResponseWriter, code int) {
	data := make(map[string]string)
	data["title"] = fmt.Sprintf("%d %s", code, http.StatusText(code))
	rd.HTML(w, code, fmt.Sprintf("status/%d", code), data, render.HTMLOptions{Layout: "layouts/status"})
}
