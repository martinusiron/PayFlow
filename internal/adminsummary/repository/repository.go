package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/adminsummary/domain"
)

type AdminSummaryRepository interface {
	GetSummaryByPayrollID(ctx context.Context, payrollID int) (*domain.FullSummary, error)
}
