package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/abspayd/music-guide/handlers"
)

func SetupRoutes(e *echo.Echo) {
	// Home page
	e.GET("/", handlers.GetHome)

	// Intervals
	e.GET("/intervals", handlers.GetIntervalCalculator)
	e.POST("/intervals", handlers.PostIntervalCalculator)

	// Circle of fifths
	e.GET("/circle", handlers.GetCircleOfFifths)
}
