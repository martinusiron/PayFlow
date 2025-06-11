package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/auth/domain"
)

type AuthRepository interface {
	GetPasswordHash(ctx context.Context, username string) (string, error)
	SaveToken(ctx context.Context, username string, token domain.Token) error
}
