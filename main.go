package main

import (
	"net/http"
	"os"

	"ngo-reporting-backend/cmd/server"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ginHandler := server.CreateGinHandler()
		ginHandler.ServeHTTP(w, r)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
