package handler

import (
	"net/http"

	"ngo-reporting-backend/cmd"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Run the main Gin server
	ginHandler := cmd.CreateGinHandler()
	ginHandler.ServeHTTP(w, r)
}
