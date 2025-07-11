package models

type Livestock struct {
	ID     int64  `json:"id"`
	FarmID int64  `json:"farm_id"`
	Type   string `json:"type"` // e.g., "Cattle", "Goats"
	Count  int    `json:"count"`
}
