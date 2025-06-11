package repository

import (
	"context"
	"database/sql"

	"github.com/martinusiron/PayFlow/internal/auth/domain"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetPasswordHash(ctx context.Context, username string) (string, error) {
	var hash string
	err := r.db.QueryRowContext(ctx,
		`SELECT password FROM users WHERE username = $1`, username).Scan(&hash)
	return hash, err
}

func (r *authRepository) SaveToken(ctx context.Context, username string, token domain.Token) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE users SET password = $1 WHERE username = $2`, token.AccessToken, username)
	return err
}
