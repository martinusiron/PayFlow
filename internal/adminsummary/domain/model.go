package domain

type AdminPayslipSummary struct {
	UserID        int
	EmployeeName  string
	TotalTakeHome float64
}

type FullSummary struct {
	PeriodID       int                   `json:"period_id"`
	TotalEmployees int                   `json:"total_employees"`
	TotalPayout    float64               `json:"total_payout"`
	Details        []AdminPayslipSummary `json:"details"`
}
