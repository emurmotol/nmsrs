package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	ctrl "github.com/zneyrl/nmsrs-lookup/app/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))
	e.GET("/", ctrl.ShowHomePage)
	e.GET("/login", ctrl.ShowLoginForm)
	e.Logger.Fatal(e.Start(":1323"))
}
