package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
)

type AuditLogService interface {
	Record(ctx context.Context, log domain.AuditLog) error
}
