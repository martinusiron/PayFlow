package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/auditlog/domain"
)

type auditLogRepo struct {
	db *sql.DB
}

func NewAuditLogRepository(db *sql.DB) *auditLogRepo {
	return &auditLogRepo{db}
}

func (r *auditLogRepo) LogAction(ctx context.Context, log domain.AuditLog) error {
	_, err := r.db.Exec(`
		INSERT INTO audit_logs (table_name, action, record_id, user_id, ip_address, request_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`, log.TableName, log.Action, log.RecordID, log.UserID, log.IPAddress, log.RequestID)
	return err
}

func (r *auditLogRepo) GetLogsByUser(ctx context.Context, userID int) ([]domain.AuditLog, error) {
	rows, err := r.db.Query(`
		SELECT id, table_name, action, record_id, user_id, ip_address, request_id, created_at
		FROM audit_logs WHERE user_id = $1
		ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []domain.AuditLog
	for rows.Next() {
		var log domain.AuditLog
		if err := rows.Scan(&log.ID, &log.TableName, &log.Action, &log.RecordID, &log.UserID, &log.IPAddress, &log.RequestID, &log.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
