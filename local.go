package main

import (
	"log"
	"net/http"
	"os"

	"ngo-reporting-backend/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
