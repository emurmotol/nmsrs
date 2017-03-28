package routes

import (
	"github.com/labstack/echo"
	"github.com/zneyrl/nmsrs-lookup/controllers"
)

func API(e *echo.Echo) *echo.Echo {
	e.GET("/", controllers.ShowHomePage)

	return e
} // TODO
