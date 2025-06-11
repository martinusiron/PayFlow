package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/payslip/domain"
)

type PayslipRepository interface {
	GetLatestPayslip(ctx context.Context, userID int) (*domain.Payslip, error)
}
