package domain

type AdminPayslipSummary struct {
	UserID        int
	EmployeeName  string
	TotalTakeHome float64
}

type FullSummary struct {
	PeriodID       int
	TotalEmployees int
	TotalPayout    float64
	Details        []AdminPayslipSummary
}
