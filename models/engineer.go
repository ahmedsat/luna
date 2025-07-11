package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Engineer{})
}

type Engineer struct {
	ID      int64    `gorm:"primaryKey" json:"id"`
	Name    string   `gorm:"not null" json:"name"`
	Regions []string `gorm:"type:text[];not null;default:'{}'" json:"regions"` // PostgreSQL text[]
}
