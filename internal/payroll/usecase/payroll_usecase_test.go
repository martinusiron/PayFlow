package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/martinusiron/PayFlow/internal/payroll/domain"
	sm "github.com/martinusiron/PayFlow/internal/shared/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRunPayroll_Success(t *testing.T) {
	repo := new(mocks.PayrollRepository)
	sharedSvc := new(sm.ServiceInterface)
	uc := NewPayrollUsecase(repo, sharedSvc)
	ctx := context.Background()

	start := "2025-06-01"
	end := "2025-06-30"
	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)

	repo.On("IsPayrollRun", ctx, start, end).Return(false, nil)
	repo.On("CreatePayroll", ctx, mock.Anything).Return(1, nil)
	sharedSvc.On("CalculateAllEmployees", ctx, startDate, endDate).Return([]domain.ProcessedPayroll{
		{
			UserID:          1,
			BaseSalary:      5000000,
			WorkdaysPresent: 20,
			ProratedSalary:  4761904,
			OvertimeHours:   5,
			OvertimePay:     476190,
			Reimbursements:  200000,
			TotalTakeHome:   5438094,
		},
	}, nil)
	repo.On("MarkAsProcessed", ctx, 1, mock.Anything).Return(nil)

	err := uc.RunPayroll(ctx, start, end, 1, "127.0.0.1", "req-xyz")
	assert.NoError(t, err)
}

func TestRunPayroll_Duplicate(t *testing.T) {
	repo := new(mocks.PayrollRepository)
	sharedSvc := new(sm.ServiceInterface)
	uc := NewPayrollUsecase(repo, sharedSvc)
	ctx := context.Background()

	repo.On("IsPayrollRun", ctx, "2025-06-01", "2025-06-30").Return(true, nil)

	err := uc.RunPayroll(ctx, "2025-06-01", "2025-06-30", 1, "ip", "req")
	assert.EqualError(t, err, "payroll for this period already run")
}
