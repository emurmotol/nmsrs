package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

func ShowLoginForm(c echo.Context) error {
	return c.String(http.StatusOK, "Login GET")
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login POST")
}
