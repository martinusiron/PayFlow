package domain

import "time"

type Payroll struct {
	ID          int
	PeriodStart time.Time
	PeriodEnd   time.Time
	RunAt       time.Time
	CreatedBy   int
	CreatedAt   time.Time
	IPAddress   string
	RequestID   string
}

type ProcessedPayroll struct {
	PayrollID       int
	UserID          int
	BaseSalary      float64
	WorkdaysPresent int
	ProratedSalary  float64
	OvertimeHours   float64
	OvertimePay     float64
	Reimbursements  float64
	TotalTakeHome   float64
}
