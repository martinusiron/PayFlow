package domain

import "time"

type Overtime struct {
	ID        int
	UserID    int
	Date      time.Time
	Hours     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int
	UpdatedBy int
	IPAddress string
	RequestID string
}
