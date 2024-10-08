package main

import (
	"os"

	"github.com/abspayd/music-guide/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "views/static")

	routes.SetupRoutes(e)

	godotenv.Load()


	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
