package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zneyrl/nmsrs-lookup/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))

	e = routes.Web(e)
	e.Logger.Fatal(e.Start(":1323"))
}
