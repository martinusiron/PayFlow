// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/martinusiron/PayFlow/internal/reimbursement/domain"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ReimbursementRepository is an autogenerated mock type for the ReimbursementRepository type
type ReimbursementRepository struct {
	mock.Mock
}

// GetReimbursementsByUser provides a mock function with given fields: ctx, userID, start, end
func (_m *ReimbursementRepository) GetReimbursementsByUser(ctx context.Context, userID int, start time.Time, end time.Time) ([]domain.Reimbursement, error) {
	ret := _m.Called(ctx, userID, start, end)

	if len(ret) == 0 {
		panic("no return value specified for GetReimbursementsByUser")
	}

	var r0 []domain.Reimbursement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time, time.Time) ([]domain.Reimbursement, error)); ok {
		return rf(ctx, userID, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time, time.Time) []domain.Reimbursement); ok {
		r0 = rf(ctx, userID, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Reimbursement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, time.Time, time.Time) error); ok {
		r1 = rf(ctx, userID, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitReimbursement provides a mock function with given fields: ctx, r
func (_m *ReimbursementRepository) SubmitReimbursement(ctx context.Context, r domain.Reimbursement) (int, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for SubmitReimbursement")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Reimbursement) (int, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Reimbursement) int); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Reimbursement) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReimbursementRepository creates a new instance of ReimbursementRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReimbursementRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReimbursementRepository {
	mock := &ReimbursementRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
