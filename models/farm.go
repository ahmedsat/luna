package models

import "time"

type Farm struct {
	ID                int64     `json:"id"` // id is the code of the farm `EG/ID`
	ArabicName        string    `json:"arabic_name"`
	EnglishName       string    `json:"english_name"`
	EngineerID        int64     `json:"engineer_id"`
	ManagerID         int64     `json:"manager_id"`
	Region            string    `json:"region"`
	TotalArea         float64   `json:"total_area"` // saved in m2
	CultivatedArea    float64   `json:"cultivated_area"`
	YearOfReclamation int       `json:"year_of_reclamation"`
	OwnershipDocument string    `json:"ownership_document"` // file path or URL
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
