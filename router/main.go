package router

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/home", makeHandlerWithTemplate(handleIndex))

	http.HandleFunc("/intervals", makeHandlerWithTemplate(handleIntervals))
	http.HandleFunc("/intervals/validation/", makeHandler(handleIntervalsValidation))
	http.HandleFunc("/intervals/octaveModeToggle", makeHandler(handleOctaveMode))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
