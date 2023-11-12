package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abspayd/music-companion/router"
)

func main() {
	port := 3000
	fmt.Printf("Starting server on port %d.\n", port)
	router.HandleRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
