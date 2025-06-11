package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/martinusiron/PayFlow/internal/adminsummary/domain"
	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSummary_Success(t *testing.T) {
	repo := new(mocks.AdminSummaryRepository)
	uc := NewAdminSummaryUsecase(repo)
	ctx := context.Background()

	expected := &domain.FullSummary{
		PeriodID:       1,
		TotalEmployees: 2,
		TotalPayout:    10000000,
		Details: []domain.AdminPayslipSummary{
			{UserID: 1, EmployeeName: "martin", TotalTakeHome: 5000000},
			{UserID: 2, EmployeeName: "sijabat", TotalTakeHome: 5000000},
		},
	}

	repo.On("GetSummaryByPayrollID", ctx, 1).Return(expected, nil)

	result, err := uc.GenerateSummary(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, result.TotalEmployees)
	assert.Equal(t, 10000000.0, result.TotalPayout)
}

func TestGenerateSummary_Failure(t *testing.T) {
	repo := new(mocks.AdminSummaryRepository)
	uc := NewAdminSummaryUsecase(repo)
	ctx := context.Background()

	repo.On("GetSummaryByPayrollID", ctx, 99).Return(nil, errors.New("not found"))

	_, err := uc.GenerateSummary(ctx, 99)
	assert.Error(t, err)
}
