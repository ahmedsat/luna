package models

type Worker struct {
	ID       int64  `json:"id"`
	FarmID   int64  `json:"farm_id"`
	AgeGroup string `json:"age_group"` // "<16","16-24", "25-40", "40-60", "60+"
	Gender   string `json:"gender"`    // "Male", "Female"
	Count    int    `json:"count"`
}
