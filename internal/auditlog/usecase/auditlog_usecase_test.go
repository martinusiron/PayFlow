package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
	"github.com/martinusiron/PayFlow/internal/auditlog/usecase"
	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRecord_Success(t *testing.T) {
	repo := new(mocks.AuditLogRepository)
	uc := usecase.NewAuditLogUsecase(repo)
	ctx := context.Background()

	log := domain.AuditLog{
		TableName: "attendances",
		Action:    "INSERT",
		RecordID:  1,
		UserID:    123,
		IPAddress: "127.0.0.1",
		RequestID: "req-abc",
	}

	repo.On("LogAction", ctx, log).Return(nil)
	err := uc.Record(ctx, log)
	assert.NoError(t, err)
}

func TestRecord_Failure(t *testing.T) {
	repo := new(mocks.AuditLogRepository)
	uc := usecase.NewAuditLogUsecase(repo)
	ctx := context.Background()

	log := domain.AuditLog{
		TableName: "attendances",
		Action:    "INSERT",
		RecordID:  1,
		UserID:    123,
		IPAddress: "127.0.0.1",
		RequestID: "req-abc",
	}

	repo.On("LogAction", ctx, log).Return(errors.New("DB error"))
	err := uc.Record(ctx, log)
	assert.Error(t, err)
}
