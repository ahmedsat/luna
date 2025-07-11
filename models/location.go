package models

type Location struct {
	ID        int64  `json:"id"`
	FarmID    int64  `json:"farm_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
