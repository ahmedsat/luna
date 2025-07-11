package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Attachment{})
}

type Attachment struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	FarmID   int64  `gorm:"not null;index" json:"farm_id"`
	Name     string `gorm:"not null" json:"name"`
	FilePath string `gorm:"not null" json:"file_path"`
}
