package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/abspayd/music-guide/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {

	// Main route
	e.GET("/", handlers.GetMain)

	e.GET("/intervals", handlers.GetIntervalCalculator)
	e.POST("/intervals", handlers.PostIntervalCalculator)
}
