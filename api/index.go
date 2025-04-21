package handler

import (
	"log"
	"net/http"
	"os"

	"ngo-reporting-backend/cmd"

	"github.com/vercel/go-bridge/go/bridge"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Run the main Gin server
	ginHandler := cmd.CreateGinHandler()
	ginHandler.ServeHTTP(w, r)
}

func main() {
	log.Println("Starting serverless function")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(bridge.Start(":" + port))
}
