package routes

import (
	"github.com/labstack/echo"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
)

func API(e *echo.Echo) *echo.Echo {
	e.GET("/", home.Index)

	return e
} // TODO
