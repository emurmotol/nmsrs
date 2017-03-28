package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zneyrl/nmsrs-lookup/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))
	e.GET("/", controllers.ShowHomePage)
	e.GET("/login", controllers.ShowLoginForm)
	e.Logger.Fatal(e.Start(":1323"))
}
