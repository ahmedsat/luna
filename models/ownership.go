package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Ownership{})
}

type Ownership struct {
	FarmID    int64   `gorm:"primaryKey;not null" json:"farm_id"`
	FarmerID  int64   `gorm:"primaryKey;not null" json:"farmer_id"`
	Share     float64 `gorm:"not null;check:share >= 0 AND share <= 100" json:"share"`
	IsManager bool    `gorm:"not null;default:false" json:"is_manager"`
}
