package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/payroll/domain"
)

type payrollRepository struct {
	db *sql.DB
}

func NewPayrollRepository(db *sql.DB) *payrollRepository {
	return &payrollRepository{db}
}

func (r *payrollRepository) CreatePayroll(ctx context.Context, p domain.Payroll) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO payroll_periods (start_date, end_date, run_at, created_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		p.PeriodStart, p.PeriodEnd, p.RunAt, p.CreatedBy, p.IPAddress, p.RequestID,
	).Scan(&id)

	return id, err
}

func (r *payrollRepository) IsPayrollRun(ctx context.Context, start, end string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM payroll_periods WHERE start_date = $1 AND end_date = $2
		)`, start, end).Scan(&exists)
	return exists, err
}

func (r *payrollRepository) MarkAsProcessed(
	ctx context.Context,
	payrollID int,
	details []domain.ProcessedPayroll,
	adminID int, ip, reqID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO payrolls
		(payroll_period_id, user_id, base_salary, workdays_present, prorated_salary, overtime_hours, overtime_amount, reimbursement_amount, total_amount, created_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range details {
		_, err := stmt.Exec(
			payrollID, d.UserID, d.BaseSalary, d.WorkdaysPresent,
			d.ProratedSalary, d.OvertimeHours, d.OvertimePay,
			d.Reimbursements, d.TotalTakeHome, adminID, ip, reqID,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
