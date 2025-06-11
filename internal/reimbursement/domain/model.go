package domain

import "time"

type Reimbursement struct {
	ID          int
	UserID      int
	Date        time.Time
	Amount      float64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   int
	UpdatedBy   int
	IPAddress   string
	RequestID   string
}
