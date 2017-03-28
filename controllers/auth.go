package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowLoginForm(c echo.Context) error {
	return c.String(http.StatusOK, "Login Form")
}
