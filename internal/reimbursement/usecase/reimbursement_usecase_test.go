package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/martinusiron/PayFlow/internal/reimbursement/domain"
	"github.com/stretchr/testify/assert"
)

func TestSubmitReimbursement_Success(t *testing.T) {
	mockRepo := new(mocks.ReimbursementRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewReimbursementUsecase(mockRepo, mockAudit)

	date := time.Now()
	amount := 150.75
	desc := "Office supplies"
	userID := 1
	ip := "127.0.0.1"
	reqID := "req-abc"

	mockRepo.On("SubmitReimbursement", context.Background(), domain.Reimbursement{
		UserID:      userID,
		Date:        date,
		Amount:      amount,
		Description: desc,
		CreatedBy:   userID,
		IPAddress:   ip,
		RequestID:   reqID,
	}).Return(nil)

	err := uc.Submit(context.Background(), userID, date, amount, desc, userID, ip, reqID)
	assert.NoError(t, err)
}

func TestSubmitReimbursement_ZeroAmount(t *testing.T) {
	mockRepo := new(mocks.ReimbursementRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewReimbursementUsecase(mockRepo, mockAudit)

	err := uc.Submit(context.Background(), 1, time.Now(), 0, "desc", 1, "127.0.0.1", "req-xyz")
	assert.EqualError(t, err, "amount must be greater than zero")
}
