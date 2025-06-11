package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
)

type AuditLogService interface {
	Record(ctx context.Context, log domain.AuditLog) error
	FetchUserLogs(ctx context.Context, userID int) ([]domain.AuditLog, error)
}
