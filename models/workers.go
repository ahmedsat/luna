package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Worker{})
}

type Worker struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	FarmID   int64  `gorm:"not null;index" json:"farm_id"`
	AgeGroup string `gorm:"not null" json:"age_group"` // "<16","16-24", etc.
	Gender   string `gorm:"type:VARCHAR(10);not null;check:gender IN ('Male', 'Female')" json:"gender"`
	Count    int    `gorm:"not null;check:count >= 0" json:"count"`
}
