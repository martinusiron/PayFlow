package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/martinusiron/PayFlow/internal/overtime/domain"
	"github.com/stretchr/testify/assert"
)

func TestSubmitOvertime_Success(t *testing.T) {
	mockRepo := new(mocks.OvertimeRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewOvertimeUsecase(mockRepo, mockAudit)
	ctx := context.Background()

	date := time.Now().Add(-24 * time.Hour)
	hours := 2.0
	userID := 1
	ip := "127.0.0.1"
	reqID := "req-123"

	mockRepo.On("SubmitOvertime", ctx, domain.Overtime{
		UserID:    userID,
		Date:      date,
		Hours:     hours,
		CreatedBy: userID,
		IPAddress: ip,
		RequestID: reqID,
	}).Return(nil)

	err := uc.Submit(ctx, userID, date, hours, userID, ip, reqID)
	assert.NoError(t, err)
}

func TestSubmitOvertime_TooMuchHours(t *testing.T) {
	mockRepo := new(mocks.OvertimeRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewOvertimeUsecase(mockRepo, mockAudit)

	date := time.Now().Add(-24 * time.Hour)
	err := uc.Submit(context.Background(), 1, date, 4.5, 1, "127.0.0.1", "req-xyz")
	assert.EqualError(t, err, "overtime cannot exceed 3 hours per day")
}
