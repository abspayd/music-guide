package main

import (
	"github.com/abspayd/music-guide/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "views/static")

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
