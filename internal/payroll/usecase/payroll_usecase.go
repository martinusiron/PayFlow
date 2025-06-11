package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/martinusiron/PayFlow/internal/payroll/domain"
	"github.com/martinusiron/PayFlow/internal/payroll/repository"
	"github.com/martinusiron/PayFlow/internal/shared"
)

type PayrollUsecase struct {
	repo          repository.PayrollRepository
	sharedService shared.ServiceInterface
}

func NewPayrollUsecase(repo repository.PayrollRepository, shared shared.ServiceInterface) *PayrollUsecase {
	return &PayrollUsecase{repo, shared}
}

func (u *PayrollUsecase) RunPayroll(ctx context.Context, start, end string, adminID int, ip, reqID string) error {
	isRun, err := u.repo.IsPayrollRun(ctx, start, end)
	if err != nil {
		return err
	}
	if isRun {
		return errors.New("payroll for this period already run")
	}

	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)

	payroll := domain.Payroll{
		PeriodStart: startDate,
		PeriodEnd:   endDate,
		RunAt:       time.Now(),
		CreatedBy:   adminID,
		IPAddress:   ip,
		RequestID:   reqID,
	}
	payrollID, err := u.repo.CreatePayroll(ctx, payroll)
	if err != nil {
		return err
	}

	records, err := u.sharedService.CalculateAllEmployees(ctx, startDate, endDate)
	if err != nil {
		return err
	}

	for i := range records {
		records[i].PayrollID = payrollID
	}

	return u.repo.MarkAsProcessed(ctx, payrollID, records)
}
