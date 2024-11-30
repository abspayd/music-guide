package main

import (
	"log"
	"os"

	"github.com/abspayd/music-guide/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")

	routes.SetupRoutes(e)

	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file: %v", err)
	}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
