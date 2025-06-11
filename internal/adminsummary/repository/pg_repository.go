package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/adminsummary/domain"
)

type adminSummaryRepository struct {
	db *sql.DB
}

func NewAdminSummaryRepository(db *sql.DB) *adminSummaryRepository {
	return &adminSummaryRepository{db}
}

func (r *adminSummaryRepository) GetSummaryByPayrollID(ctx context.Context, payrollID int) (*domain.FullSummary, error) {
	rows, err := r.db.Query(`
		SELECT u.id, u.name, p.total_take_home
		FROM processed_payrolls p
		JOIN users u ON p.user_id = u.id
		WHERE p.payroll_id = $1`, payrollID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var totalPayout float64
	var summaries []domain.AdminPayslipSummary

	for rows.Next() {
		var summary domain.AdminPayslipSummary
		err := rows.Scan(&summary.UserID, &summary.EmployeeName, &summary.TotalTakeHome)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, summary)
		totalPayout += summary.TotalTakeHome
	}

	return &domain.FullSummary{
		PeriodID:       payrollID,
		TotalEmployees: len(summaries),
		TotalPayout:    totalPayout,
		Details:        summaries,
	}, nil
}
