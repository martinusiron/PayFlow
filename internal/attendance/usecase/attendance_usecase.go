package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/domain"
	"github.com/martinusiron/PayFlow/internal/attendance/repository"
	au "github.com/martinusiron/PayFlow/internal/auditlog/usecase"
)

type AttendanceUsecase struct {
	repo         repository.AttendanceRepository
	AuditUsecase au.AuditLogService
}

func NewAttendanceUsecase(

	r repository.AttendanceRepository, au au.AuditLogService) *AttendanceUsecase {
	return &AttendanceUsecase{r, au}
}

func (u *AttendanceUsecase) Submit(ctx context.Context, userID int, date time.Time, createdBy int, ip, reqID string) error {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return errors.New("attendance cannot be submitted on weekends")
	}
	att := domain.Attendance{
		UserID:    userID,
		Date:      date,
		CreatedBy: createdBy,
		IPAddress: ip,
		RequestID: reqID,
	}
	return u.repo.SubmitAttendance(ctx, att)
}
