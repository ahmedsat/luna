package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Location{})
}

type Location struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	FarmID    int64  `gorm:"not null;index" json:"farm_id"`
	Latitude  string `gorm:"not null" json:"latitude"`
	Longitude string `gorm:"not null" json:"longitude"`
}
