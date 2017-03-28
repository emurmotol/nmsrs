package routes

import (
	"github.com/labstack/echo"
	"github.com/zneyrl/nmsrs-lookup/controllers"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
)

func Web(e *echo.Echo) *echo.Echo {
	e.GET("/", controllers.ShowHomePage)

	e.GET("/login", auth.ShowLoginForm)
	e.POST("/login", auth.Login)

	return e
}
