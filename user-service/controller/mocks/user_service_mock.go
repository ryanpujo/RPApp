package mocks

import (
	"context"

	"github.com/spriigan/RPApp/repository"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	args := m.Called(ctx, arg)
	user, _ := args.Get(0).(repository.User)
	return user, args.Error(1)
}

func (m *MockUserService) DeleteByID(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) GetById(ctx context.Context, id int64) (repository.User, error) {
	args := m.Called(ctx, id)
	user, _ := args.Get(0).(repository.User)
	return user, args.Error(1)
}

func (m *MockUserService) GetMany(ctx context.Context, limit, page int32) ([]repository.User, error) {
	args := m.Called(ctx, limit)
	users, _ := args.Get(0).([]repository.User)
	return users, args.Error(1)
}

func (m *MockUserService) UpdateByID(ctx context.Context, arg repository.UpdateByIDParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}
