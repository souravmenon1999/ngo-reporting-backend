package controllers

import (
	"net/http"
	"time"

	"ngo-reporting-backend/models"
	"ngo-reporting-backend/services"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service *services.ReportService
}

func NewReportController(service *services.ReportService) *ReportController {
	return &ReportController{service: service}
}

func (c *ReportController) SubmitReport(ctx *gin.Context) {
	var report models.Report
	if err := ctx.ShouldBindJSON(&report); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Basic validation
	if report.NGOID == "" || report.Month == "" || report.PeopleHelped < 0 || report.EventsConducted < 0 || report.FundsUtilized < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report data"})
		return
	}

	// Parse and validate month format (YYYY-MM)
	_, err := time.Parse("2006-01", report.Month)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month format, use YYYY-MM"})
		return
	}

	if err := c.service.SaveReport(&report); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save report"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Report submitted successfully"})
}

func (c *ReportController) GetDashboard(ctx *gin.Context) {
	month := ctx.Query("month")
	if month == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required"})
		return
	}

	// Validate month format
	_, err := time.Parse("2006-01", month)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month format, use YYYY-MM"})
		return
	}

	dashboard, err := c.service.GetDashboardData(month)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch dashboard data"})
		return
	}

	ctx.JSON(http.StatusOK, dashboard)
}