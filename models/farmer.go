package models

import "github.com/ahmedsat/luna/db"

func init() {
	db.MigrationTable = append(db.MigrationTable, &Farmer{})
}

type Farmer struct {
	ID            int    `gorm:"primaryKey;unique;check:id >= 10000000000000 AND id <= 99999999999999" json:"id"` // 14-digit ID
	FullName      string `gorm:"not null" json:"full_name"`
	Gender        string `gorm:"type:VARCHAR(10);not null;check:gender IN ('Male', 'Female')" json:"gender"`
	PersonalImage string `json:"personal_image"`
	IDImageFront  string `json:"id_image_front"`
	IDImageBack   string `json:"id_image_back"`
	PhoneNumber   string `gorm:"type:VARCHAR(20)" json:"phone_number"`
}
