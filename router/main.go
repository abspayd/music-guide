package router

import (
	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {

	e.GET("/", handleMain)

	e.Static("/static", "assets")

	// -- deprecated --
	// http.HandleFunc("/", handleDefault)
	// http.HandleFunc("/home", makeHandler(handleIndex))
	// http.HandleFunc("/intervals", makeHandler(handleIntervals))
	// http.HandleFunc("/circle-quiz", makeHandler(handleCircleOfFifthsQuiz))

	// http.HandleFunc("/intervalsValidation/", handleIntervalsValidation)
	// http.HandleFunc("/intervalsOctaveModeToggle", handleOctaveMode)

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
