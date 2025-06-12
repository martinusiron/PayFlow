package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/domain"
	ald "github.com/martinusiron/PayFlow/internal/auditlog/domain"
	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitAttendance(t *testing.T) {
	mockRepo := new(mocks.AttendanceRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewAttendanceUsecase(mockRepo, mockAudit)
	ctx := context.Background()

	date := time.Date(2025, 6, 10, 0, 0, 0, 0, time.UTC)
	userID := 1
	ip := "127.0.0.1"
	reqID := "req-123"

	mockRepo.On("SubmitAttendance", ctx, domain.Attendance{
		UserID:    userID,
		Date:      date,
		CreatedBy: userID,
		IPAddress: ip,
		RequestID: reqID,
	}).Return(1, nil)

	mockAudit.On("Record", ctx, mock.MatchedBy(func(log ald.AuditLog) bool {
		return log.TableName == "attendances" &&
			log.Action == "submit" &&
			log.RecordID == 1 &&
			log.UserID == userID &&
			log.IPAddress == ip &&
			log.RequestID == reqID
	})).Return(nil)

	err := uc.Submit(ctx, userID, date, userID, ip, reqID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
	mockAudit.AssertExpectations(t)
}

func TestWeekendAttendance(t *testing.T) {
	mockRepo := new(mocks.AttendanceRepository)
	mockAudit := new(mocks.AuditLogService)
	uc := NewAttendanceUsecase(mockRepo, mockAudit)

	date := time.Date(2025, 6, 8, 0, 0, 0, 0, time.UTC) // Sunday
	err := uc.Submit(context.Background(), 1, date, 1, "127.0.0.1", "req-456")
	assert.EqualError(t, err, "attendance cannot be submitted on weekends")
}
