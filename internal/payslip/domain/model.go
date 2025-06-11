package domain

type Payslip struct {
	UserID          int
	BaseSalary      float64
	WorkdaysPresent int
	ProratedSalary  float64
	OvertimeHours   float64
	OvertimePay     float64
	Reimbursements  float64
	TotalTakeHome   float64
}
