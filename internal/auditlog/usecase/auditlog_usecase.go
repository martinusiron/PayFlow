package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
	"github.com/martinusiron/PayFlow/internal/auditlog/repository"
)

type AuditLogUsecase struct {
	repo repository.AuditLogRepository
}

func NewAuditLogUsecase(repo repository.AuditLogRepository) *AuditLogUsecase {
	return &AuditLogUsecase{repo}
}

func (u *AuditLogUsecase) Record(ctx context.Context, log domain.AuditLog) error {
	return u.repo.LogAction(ctx, log)
}
