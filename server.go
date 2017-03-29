package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zneyrl/nmsrs-lookup/routes"
	"github.com/zneyrl/nmsrs-lookup/shared/templates"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Static("static"))

	e.Renderer = &templates.Template{}
	templates.Load()

	e = routes.Web(e) // TODO
	e.Logger.Fatal(e.Start(":1323"))
}
