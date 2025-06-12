package shared

import (
	"context"
	"time"

	pd "github.com/martinusiron/PayFlow/internal/payroll/domain"
)

type ServiceInterface interface {
	CalculateAllEmployees(ctx context.Context, start, end time.Time) ([]pd.ProcessedPayroll, error)
}
