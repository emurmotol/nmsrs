package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowHomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Home GET")
}
