package models

type Ownership struct {
	FarmID    int64   `json:"farm_id"`
	FarmerID  int64   `json:"farmer_id"`
	Share     float64 `json:"share"`
	IsManager bool    `json:"is_manager"` // only one per farm
}
