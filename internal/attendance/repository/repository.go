package repository

import (
	"context"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/domain"
)

type AttendanceRepository interface {
	SubmitAttendance(ctx context.Context, att domain.Attendance) error
	GetAttendanceByUser(ctx context.Context, userID int, start, end string) ([]domain.Attendance, error)
	CountWeekdaysByUserID(ctx context.Context, userID int, start, end time.Time) (int, error)
}
