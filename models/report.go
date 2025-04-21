package models

type Report struct {
	ID              string  `json:"id" bson:"_id,omitempty"`
	NGOID           string  `json:"ngo_id" bson:"ngo_id"`
	Month           string  `json:"month" bson:"month"` // Format: YYYY-MM
	PeopleHelped    int     `json:"people_helped" bson:"people_helped"`
	EventsConducted int     `json:"events_conducted" bson:"events_conducted"`
	FundsUtilized   float64 `json:"funds_utilized" bson:"funds_utilized"`
}

type DashboardData struct {
	TotalNGOs   int     `bson:"total_ngos" json:"total_ngos"`
	TotalPeople int     `bson:"total_people" json:"total_people"`
	TotalEvents int     `bson:"total_events" json:"total_events"`
	TotalFunds  float64 `bson:"total_funds" json:"total_funds"`
}
