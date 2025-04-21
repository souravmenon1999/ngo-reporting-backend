package main

import (
	"log"
	"net/http"
	"os"

	"ngo-reporting-backend/config"
	"ngo-reporting-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Handler is the exported function for Vercel serverless
func Handler(w http.ResponseWriter, r *http.Request) {
	router := setupRouter()
	router.ServeHTTP(w, r)
}

func setupRouter() *gin.Engine {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize MongoDB
	if err := config.InitDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Add CORS middleware
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{frontendURL}
	configCors.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	configCors.AllowHeaders = []string{"Origin", "Content-Type"}
	configCors.AllowCredentials = true
	router.Use(cors.New(configCors))

	// Initialize routes
	routes.SetupRoutes(router)

	return router
}

func main() {
	// For local development
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
