package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinusiron/PayFlow/internal/attendance/domain"
)

type attendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *attendanceRepository {
	return &attendanceRepository{db}
}

func (r *attendanceRepository) SubmitAttendance(ctx context.Context, att domain.Attendance) error {
	_, err := r.db.Exec(`
		INSERT INTO attendance (user_id, attendance_date, created_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id, attendance_date) DO NOTHING`,
		att.UserID, att.Date, att.CreatedBy, att.IPAddress, att.RequestID)
	return err
}

func (r *attendanceRepository) GetAttendanceByUser(ctx context.Context, userID int, start, end string) ([]domain.Attendance, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, attendance_date, created_at, updated_at
		FROM attendance
		WHERE user_id = $1 AND attendance_date BETWEEN $2 AND $3`, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []domain.Attendance
	for rows.Next() {
		var a domain.Attendance
		err := rows.Scan(&a.ID, &a.UserID, &a.Date, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, a)
	}
	return records, nil
}

func (r *attendanceRepository) CountWeekdaysByUserID(ctx context.Context, userID int, start, end time.Time) (int, error) {
	query := `
		SELECT COUNT(DISTINCT attendance_date)
		FROM attendance
		WHERE user_id = $1
		AND attendance_date BETWEEN $2 AND $3
		AND EXTRACT(DOW FROM attendance_date) BETWEEN 1 AND 5 -- Mon(1) to Fri(5)
	`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID, start, end).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
