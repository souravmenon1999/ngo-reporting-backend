package services

import (
	"context"
	"log"
	"time"

	"ngo-reporting-backend/config"
	"ngo-reporting-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReportService struct {
	collection *mongo.Collection
}

func NewReportService() *ReportService {
	collection := config.DB.Collection("reports")
	return &ReportService{collection: collection}
}

func (s *ReportService) SaveReport(report *models.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, report)
	return err
}

func (s *ReportService) GetDashboardData(month string) (*models.DashboardData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Querying dashboard data for month: %s", month)

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"month": month}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":          nil,
			"total_ngos":   bson.M{"$addToSet": "$ngo_id"},
			"total_people": bson.M{"$sum": "$people_helped"},
			"total_events": bson.M{"$sum": "$events_conducted"},
			"total_funds":  bson.M{"$sum": "$funds_utilized"},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"total_ngos":   bson.M{"$size": "$total_ngos"},
			"total_people": 1,
			"total_events": 1,
			"total_funds":  1,
			"_id":          0,
		}}},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Printf("Aggregation error: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Read all results into a slice of bson.M for logging
	var rawResults []bson.M
	if err = cursor.All(ctx, &rawResults); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	log.Printf("Raw aggregation results for month %s: %v", month, rawResults)

	// Convert rawResults to DashboardData
	var results []models.DashboardData
	for _, raw := range rawResults {
		var data models.DashboardData
		bsonBytes, _ := bson.Marshal(raw)
		if err := bson.Unmarshal(bsonBytes, &data); err != nil {
			log.Printf("Unmarshal error: %v", err)
			continue
		}
		results = append(results, data)
	}

	if len(results) == 0 {
		// Return empty data if no results found
		return &models.DashboardData{
			TotalNGOs:   0,
			TotalPeople: 0,
			TotalEvents: 0,
			TotalFunds:  0,
		}, nil
	}

	return &results[0], nil
}
