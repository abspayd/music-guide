package router

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/home", makeHandler(handleIndex))
	http.HandleFunc("/intervals", makeHandler(handleIntervals))

	http.HandleFunc("/validateNote/", handleValidateNote)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
