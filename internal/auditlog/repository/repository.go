package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
)

type AuditLogRepository interface {
	LogAction(ctx context.Context, log domain.AuditLog) error
}
