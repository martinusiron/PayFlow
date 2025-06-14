// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/martinusiron/PayFlow/internal/payroll/domain"
	mock "github.com/stretchr/testify/mock"
)

// PayrollRepository is an autogenerated mock type for the PayrollRepository type
type PayrollRepository struct {
	mock.Mock
}

// CreatePayroll provides a mock function with given fields: ctx, p
func (_m *PayrollRepository) CreatePayroll(ctx context.Context, p domain.Payroll) (int, error) {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayroll")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Payroll) (int, error)); ok {
		return rf(ctx, p)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Payroll) int); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Payroll) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsPayrollRun provides a mock function with given fields: ctx, start, end
func (_m *PayrollRepository) IsPayrollRun(ctx context.Context, start string, end string) (bool, error) {
	ret := _m.Called(ctx, start, end)

	if len(ret) == 0 {
		panic("no return value specified for IsPayrollRun")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, start, end)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkAsProcessed provides a mock function with given fields: ctx, payrollID, details, adminID, ip, reqID
func (_m *PayrollRepository) MarkAsProcessed(ctx context.Context, payrollID int, details []domain.ProcessedPayroll, adminID int, ip string, reqID string) error {
	ret := _m.Called(ctx, payrollID, details, adminID, ip, reqID)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsProcessed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []domain.ProcessedPayroll, int, string, string) error); ok {
		r0 = rf(ctx, payrollID, details, adminID, ip, reqID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPayrollRepository creates a new instance of PayrollRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPayrollRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PayrollRepository {
	mock := &PayrollRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
