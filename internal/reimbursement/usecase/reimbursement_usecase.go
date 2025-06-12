package usecase

import (
	"context"
	"errors"
	"time"

	ald "github.com/martinusiron/PayFlow/internal/auditlog/domain"
	au "github.com/martinusiron/PayFlow/internal/auditlog/usecase"
	"github.com/martinusiron/PayFlow/internal/reimbursement/domain"
	"github.com/martinusiron/PayFlow/internal/reimbursement/repository"
)

type ReimbursementUsecase struct {
	repo         repository.ReimbursementRepository
	AuditUsecase au.AuditLogService
}

func NewReimbursementUsecase(
	r repository.ReimbursementRepository,
	au au.AuditLogService) *ReimbursementUsecase {
	return &ReimbursementUsecase{r, au}
}

func (u *ReimbursementUsecase) Submit(ctx context.Context, userID int, date time.Time, amount float64, desc string, createdBy int, ip, reqID string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	rb := domain.Reimbursement{
		UserID:      userID,
		Date:        date,
		Amount:      amount,
		Description: desc,
		CreatedBy:   createdBy,
		IPAddress:   ip,
		RequestID:   reqID,
	}
	id, err := u.repo.SubmitReimbursement(ctx, rb)
	if err != nil {
		return err
	}

	return u.AuditUsecase.Record(ctx, ald.AuditLog{
		UserID:    userID,
		TableName: "reimbursements",
		Action:    "submit",
		RecordID:  id,
		IPAddress: ip,
		RequestID: reqID,
		CreatedAt: time.Now(),
	})
}
