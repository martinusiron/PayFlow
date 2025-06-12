package domain

type Payslip struct {
	UserID          int     `json:"user_id"`
	BaseSalary      float64 `json:"base_salary"`
	WorkdaysPresent int     `json:"workdays_present"`
	ProratedSalary  float64 `json:"prorated_salary"`
	OvertimeHours   float64 `json:"overtime_hours"`
	OvertimePay     float64 `json:"overtime_pay"`
	Reimbursements  float64 `json:"reimbursements"`
	TotalTakeHome   float64 `json:"total_take_home"`
}
