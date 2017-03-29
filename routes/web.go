package routes

import (
	"github.com/labstack/echo"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
)

func Web(e *echo.Echo) *echo.Echo {
	e.GET("/", home.Index)

	e.GET("/login", auth.ShowLoginForm)
	e.POST("/login", auth.Login)

	return e
}
