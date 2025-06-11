package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/adminsummary/domain"
	"github.com/martinusiron/PayFlow/internal/adminsummary/repository"
)

type AdminSummaryUsecase struct {
	repo repository.AdminSummaryRepository
}

func NewAdminSummaryUsecase(repo repository.AdminSummaryRepository) *AdminSummaryUsecase {
	return &AdminSummaryUsecase{repo}
}

func (u *AdminSummaryUsecase) GenerateSummary(ctx context.Context, payrollID int) (*domain.FullSummary, error) {
	return u.repo.GetSummaryByPayrollID(ctx, payrollID)
}
