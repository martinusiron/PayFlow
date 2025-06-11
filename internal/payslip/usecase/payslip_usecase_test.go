package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/martinusiron/PayFlow/internal/payslip/domain"
	"github.com/stretchr/testify/assert"
)

func TestGenerateLatestPayslip_Success(t *testing.T) {
	repo := new(mocks.PayslipRepository)
	uc := NewPayslipUsecase(repo)
	ctx := context.Background()

	expected := &domain.Payslip{
		UserID:          1,
		BaseSalary:      5000000,
		WorkdaysPresent: 20,
		ProratedSalary:  4761904,
		OvertimeHours:   5,
		OvertimePay:     476190,
		Reimbursements:  200000,
		TotalTakeHome:   5438094,
	}

	repo.On("GetLatestPayslip", ctx, 1).Return(expected, nil)

	result, err := uc.GenerateLatestPayslip(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expected.TotalTakeHome, result.TotalTakeHome)
}

func TestGenerateLatestPayslip_Error(t *testing.T) {
	repo := new(mocks.PayslipRepository)
	uc := NewPayslipUsecase(repo)
	ctx := context.Background()

	repo.On("GetLatestPayslip", ctx, 1).Return(nil, errors.New("not found"))

	_, err := uc.GenerateLatestPayslip(ctx, 1)
	assert.Error(t, err)
}
