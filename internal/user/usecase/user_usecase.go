package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/user/domain"
	"github.com/martinusiron/PayFlow/internal/user/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UserUsecase) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return u.repo.GetByUsername(ctx, username)
}

func (u *UserUsecase) GetAllEmployees(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAllEmployees(ctx)
}
