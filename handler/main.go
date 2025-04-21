package handler

import (
	"log"
	"net/http"

	"ngo-reporting-backend/config"
	"ngo-reporting-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := setupRouter()
	router.ServeHTTP(w, r)
}

func setupRouter() *gin.Engine {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := config.InitDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	router := gin.Default()

	// Log requests for debugging
	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s, Origin: %s", c.Request.Method, c.Request.URL, c.Request.Header.Get("Origin"))
		c.Next()
	})

	// CORS middleware
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true // Allow all origins in production
	configCors.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	configCors.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	configCors.ExposeHeaders = []string{"Content-Length"}
	// Note: AllowCredentials is omitted as it conflicts with AllowAllOrigins
	router.Use(cors.New(configCors))

	routes.SetupRoutes(router)
	return router
}
