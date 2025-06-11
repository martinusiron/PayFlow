package usecase

import (
	"context"

	"github.com/martinusiron/PayFlow/internal/payslip/domain"
	"github.com/martinusiron/PayFlow/internal/payslip/repository"
)

type PayslipUsecase struct {
	repo repository.PayslipRepository
}

func NewPayslipUsecase(repo repository.PayslipRepository) *PayslipUsecase {
	return &PayslipUsecase{repo}
}

func (u *PayslipUsecase) GenerateLatestPayslip(ctx context.Context, userID int) (*domain.Payslip, error) {
	return u.repo.GetLatestPayslip(ctx, userID)
}
