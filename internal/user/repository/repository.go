package repository

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/user/domain"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetAllEmployees(ctx context.Context) ([]domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	SeedIfEmpty(ctx context.Context) error
}
