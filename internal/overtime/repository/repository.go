package repository

import (
	"context"
	"time"

	"github.com/martinusiron/PayFlow/internal/overtime/domain"
)

type OvertimeRepository interface {
	SubmitOvertime(ctx context.Context, ot domain.Overtime) error
	GetOvertimeByUser(ctx context.Context, userID int, start, end time.Time) ([]domain.Overtime, error)
}
