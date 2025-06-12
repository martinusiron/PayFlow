package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/payroll/domain"
)

type PayrollRepository interface {
	CreatePayroll(ctx context.Context, p domain.Payroll) (int, error)
	MarkAsProcessed(ctx context.Context, payrollID int, details []domain.ProcessedPayroll, adminID int, ip, reqID string) error
	IsPayrollRun(ctx context.Context, start, end string) (bool, error)
}
