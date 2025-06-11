package usecase

import (
	"context"
	"errors"
	"time"

	au "github.com/martinusiron/PayFlow/internal/auditlog/usecase"
	"github.com/martinusiron/PayFlow/internal/overtime/domain"
	"github.com/martinusiron/PayFlow/internal/overtime/repository"
)

type OvertimeUsecase struct {
	repo         repository.OvertimeRepository
	AuditUsecase au.AuditLogService
}

func NewOvertimeUsecase(
	r repository.OvertimeRepository,
	au au.AuditLogService) *OvertimeUsecase {
	return &OvertimeUsecase{r, au}
}

func (u *OvertimeUsecase) Submit(ctx context.Context, userID int, date time.Time, hours float64, createdBy int, ip, reqID string) error {
	if hours > 3 {
		return errors.New("overtime cannot exceed 3 hours per day")
	}
	if date.After(time.Now()) {
		return errors.New("cannot submit overtime for future dates")
	}
	ot := domain.Overtime{
		UserID:    userID,
		Date:      date,
		Hours:     hours,
		CreatedBy: createdBy,
		IPAddress: ip,
		RequestID: reqID,
	}
	return u.repo.SubmitOvertime(ctx, ot)
}
