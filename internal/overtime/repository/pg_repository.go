package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinusiron/PayFlow/internal/overtime/domain"
)

type overtimeRepository struct {
	db *sql.DB
}

func NewOvertimeRepository(db *sql.DB) *overtimeRepository {
	return &overtimeRepository{db}
}

func (r *overtimeRepository) SubmitOvertime(ctx context.Context, ot domain.Overtime) (int, error) {
	var id int
	err := r.db.QueryRow(`
		INSERT INTO overtime (user_id, overtime_date, hours, created_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		ot.UserID, ot.Date, ot.Hours, ot.CreatedBy, ot.IPAddress, ot.RequestID).Scan(&id)
	return id, err
}

func (r *overtimeRepository) GetOvertimeByUser(ctx context.Context, userID int, start, end time.Time) ([]domain.Overtime, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, overtime_date, hours, created_at, updated_at
		FROM overtime
		WHERE user_id = $1 AND overtime_date BETWEEN $2 AND $3`, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []domain.Overtime
	for rows.Next() {
		var ot domain.Overtime
		err := rows.Scan(&ot.ID, &ot.UserID, &ot.Date, &ot.Hours, &ot.CreatedAt, &ot.UpdatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, ot)
	}
	return records, nil
}
