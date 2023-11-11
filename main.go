package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abspayd/music-companion/router"
)

const (
	PORT = 3000
)

func main() {
	fmt.Printf("Starting server on port %d...\n", PORT)
	router.HandleRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
