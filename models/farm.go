package models

import (
	"time"

	"github.com/ahmedsat/luna/db"
)

func init() {
	db.MigrationTable = append(db.MigrationTable, &Farm{})
}

type Farm struct {
	ID                int64     `gorm:"primaryKey" json:"id"` // Format is EG/{id}
	ArabicName        string    `gorm:"not null;unique" json:"arabic_name"`
	EnglishName       string    `gorm:"not null;unique" json:"english_name"`
	EngineerID        int64     `gorm:"not null;index" json:"engineer_id"`
	ManagerID         int64     `gorm:"not null;index" json:"manager_id"`
	Region            string    `gorm:"not null" json:"region"`
	TotalArea         float64   `gorm:"not null" json:"total_area"`
	CultivatedArea    float64   `gorm:"not null" json:"cultivated_area"`
	YearOfReclamation int       `json:"year_of_reclamation"`
	OwnershipDocument string    `json:"ownership_document"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
