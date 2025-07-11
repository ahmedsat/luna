package models

import (
	"time"

	"github.com/ahmedsat/luna/db"
)

func init() {
	db.MigrationTable = append(db.MigrationTable, &AuditLog{})
}

type AuditLog struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Action    string    `gorm:"type:VARCHAR(10);not null" json:"action"` // CREATE, UPDATE, DELETE
	TableName string    `gorm:"not null" json:"table_name"`
	User      string    `gorm:"not null" json:"user"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
	OldValue  string    `gorm:"type:JSONB" json:"old_value"`
	NewValue  string    `gorm:"type:JSONB" json:"new_value"`
}
