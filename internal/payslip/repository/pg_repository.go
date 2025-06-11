package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/payslip/domain"
)

type payslipRepository struct {
	db *sql.DB
}

func NewPayslipRepository(db *sql.DB) *payslipRepository {
	return &payslipRepository{db}
}

func (r *payslipRepository) GetLatestPayslip(ctx context.Context, userID int) (*domain.Payslip, error) {
	var p domain.Payslip
	err := r.db.QueryRow(`
		SELECT user_id, base_salary, workdays_present, prorated_salary,
		       overtime_hours, overtime_pay, reimbursements, total_take_home
		FROM processed_payrolls
		WHERE user_id = $1
		ORDER BY payroll_id DESC LIMIT 1`, userID).Scan(
		&p.UserID, &p.BaseSalary, &p.WorkdaysPresent, &p.ProratedSalary,
		&p.OvertimeHours, &p.OvertimePay, &p.Reimbursements, &p.TotalTakeHome,
	)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
