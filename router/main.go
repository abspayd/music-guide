package router

import (
	"net/http"
)

func HandleRoutes() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/home", makeHandler(handleIndex))
	http.HandleFunc("/intervals", makeHandler(handleIntervals))
}
