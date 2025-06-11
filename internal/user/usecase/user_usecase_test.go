package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/martinusiron/PayFlow/internal/mocks"
	"github.com/martinusiron/PayFlow/internal/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_GetByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	expectedUser := &domain.User{
		ID:       1,
		Username: "employee1",
		Salary:   10000000,
		Role:     "employee",
	}

	mockRepo.On("GetByID", mock.Anything, 1).Return(expectedUser, nil)

	user, err := uc.GetByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	mockRepo.On("GetByID", mock.Anything, 999).Return(nil, errors.New("not found"))

	user, err := uc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, user)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetByUsername(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	expectedUser := &domain.User{
		ID:       2,
		Username: "admin",
		Salary:   0,
		Role:     "admin",
	}

	mockRepo.On("GetByUsername", mock.Anything, "admin").Return(expectedUser, nil)

	user, err := uc.GetByUsername(context.Background(), "admin")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetAllEmployees(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUserUsecase(mockRepo)

	expectedUsers := []domain.User{
		{ID: 1, Username: "user1", Salary: 10000000, Role: "employee"},
		{ID: 2, Username: "user2", Salary: 8000000, Role: "employee"},
	}

	mockRepo.On("GetAllEmployees", mock.Anything).Return(expectedUsers, nil)

	users, err := uc.GetAllEmployees(context.Background())
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}
