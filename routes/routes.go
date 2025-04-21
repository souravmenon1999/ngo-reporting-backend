package routes

import (
	"ngo-reporting-backend/controllers"
	"ngo-reporting-backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	reportService := services.NewReportService()
	reportController := controllers.NewReportController(reportService)

	api := router.Group("/api")
	{
		api.POST("/report", reportController.SubmitReport)
		api.GET("/dashboard", reportController.GetDashboard)
	}
}