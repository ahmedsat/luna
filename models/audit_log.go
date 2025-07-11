package models

import "time"

type AuditLog struct {
	ID        int64     `json:"id"`
	Action    string    `json:"action"`     // CREATE, UPDATE, DELETE
	TableName string    `json:"table_name"` // "farms", "farmers", etc.
	User      string    `json:"user"`       // username or ID
	Timestamp time.Time `json:"timestamp"`
	OldValue  string    `json:"old_value"` // JSON string
	NewValue  string    `json:"new_value"` // JSON string
}
