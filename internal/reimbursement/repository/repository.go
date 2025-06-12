package repository

import (
	"context"
	"time"

	"github.com/martinusiron/PayFlow/internal/reimbursement/domain"
)

type ReimbursementRepository interface {
	SubmitReimbursement(ctx context.Context, r domain.Reimbursement) (int, error)
	GetReimbursementsByUser(ctx context.Context, userID int, start, end time.Time) ([]domain.Reimbursement, error)
}
