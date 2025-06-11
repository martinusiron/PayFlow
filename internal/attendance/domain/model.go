package domain

import "time"

type Attendance struct {
	ID         int
	UserID     int
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  int
	UpdatedBy  int
	IPAddress  string
	RequestID  string
}
