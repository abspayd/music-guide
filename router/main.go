package router

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/home", makeHandler(handleIndex))
	http.HandleFunc("/intervals", makeHandler(handleIntervals))

	http.HandleFunc("/intervalsValidation/", handleIntervalsValidation)
	http.HandleFunc("/intervalsOctaveModeToggle", handleOctaveMode)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
