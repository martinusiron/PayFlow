package shared

import (
	"context"
	"net/http"
	"time"

	pd "github.com/martinusiron/PayFlow/internal/payroll/domain"
)

type ServiceInterface interface {
	ExtractRequestContext(ctx context.Context, r *http.Request) RequestContext
	LogAudit(ctx context.Context, userID int, action, table string, recordID int, requestID, ip string)
	CalculateAllEmployees(ctx context.Context, start, end time.Time) ([]pd.ProcessedPayroll, error)
}
