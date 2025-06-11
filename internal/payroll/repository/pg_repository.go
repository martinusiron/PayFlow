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
		INSERT INTO payrolls (period_start, period_end, run_at, created_by, ip_address, request_id)
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
			SELECT 1 FROM payrolls WHERE period_start = $1 AND period_end = $2
		)`, start, end).Scan(&exists)
	return exists, err
}

func (r *payrollRepository) MarkAsProcessed(ctx context.Context, payrollID int, details []domain.ProcessedPayroll) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO processed_payrolls
		(payroll_id, user_id, base_salary, workdays_present, prorated_salary, overtime_hours, overtime_pay, reimbursements, total_take_home)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range details {
		_, err := stmt.Exec(
			d.PayrollID, d.UserID, d.BaseSalary, d.WorkdaysPresent,
			d.ProratedSalary, d.OvertimeHours, d.OvertimePay,
			d.Reimbursements, d.TotalTakeHome,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *payrollRepository) GetProcessedPayrolls(ctx context.Context, payrollID int) ([]domain.ProcessedPayroll, error) {
	rows, err := r.db.Query(`
		SELECT payroll_id, user_id, base_salary, workdays_present, prorated_salary,
		       overtime_hours, overtime_pay, reimbursements, total_take_home
		FROM processed_payrolls WHERE payroll_id = $1`, payrollID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.ProcessedPayroll
	for rows.Next() {
		var d domain.ProcessedPayroll
		err := rows.Scan(
			&d.PayrollID, &d.UserID, &d.BaseSalary, &d.WorkdaysPresent,
			&d.ProratedSalary, &d.OvertimeHours, &d.OvertimePay,
			&d.Reimbursements, &d.TotalTakeHome)
		if err != nil {
			return nil, err
		}
		results = append(results, d)
	}
	return results, nil
}
