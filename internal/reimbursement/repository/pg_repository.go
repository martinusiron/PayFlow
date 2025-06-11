package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinusiron/PayFlow/internal/reimbursement/domain"
)

type reimbursementRepository struct {
	db *sql.DB
}

func NewReimbursementRepository(db *sql.DB) *reimbursementRepository {
	return &reimbursementRepository{db}
}

func (r *reimbursementRepository) SubmitReimbursement(ctx context.Context, rb domain.Reimbursement) error {
	_, err := r.db.Exec(`
		INSERT INTO reimbursements (user_id, reimbursement_date, amount, description, created_by, ip_address, request_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		rb.UserID, rb.Date, rb.Amount, rb.Description, rb.CreatedBy, rb.IPAddress, rb.RequestID)
	return err
}

func (r *reimbursementRepository) GetReimbursementsByUser(ctx context.Context, userID int, start, end time.Time) ([]domain.Reimbursement, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, reimbursement_date, amount, description, created_at, updated_at
		FROM reimbursements
		WHERE user_id = $1 AND reimbursement_date BETWEEN $2 AND $3`, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []domain.Reimbursement
	for rows.Next() {
		var rb domain.Reimbursement
		err := rows.Scan(&rb.ID, &rb.UserID, &rb.Date, &rb.Amount, &rb.Description, &rb.CreatedAt, &rb.UpdatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, rb)
	}
	return records, nil
}
