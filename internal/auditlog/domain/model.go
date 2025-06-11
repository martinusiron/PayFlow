package domain

import "time"

type AuditLog struct {
	ID        int
	TableName string
	Action    string
	RecordID  int
	UserID    int
	IPAddress string
	RequestID string
	CreatedAt time.Time
}
