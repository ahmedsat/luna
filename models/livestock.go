package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Livestock{})
}

type Livestock struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	FarmID int64  `gorm:"not null;index" json:"farm_id"`
	Type   string `gorm:"not null" json:"type"`
	Count  int    `gorm:"not null;check:count >= 0" json:"count"`
}
