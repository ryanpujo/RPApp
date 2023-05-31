package mocks

import (
	"context"

	"github.com/spriigan/RPApp/repository"
	"github.com/stretchr/testify/mock"
)

type MockQuery struct {
	mock.Mock
}

func (m *MockQuery) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	args := m.Called(ctx, arg)
	user, _ := args.Get(0).(repository.User)
	return user, args.Error(1)
}

func (m *MockQuery) DeleteByID(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockQuery) GetById(ctx context.Context, id int64) (repository.User, error) {
	args := m.Called(ctx, id)
	user, _ := args.Get(0).(repository.User)
	return user, args.Error(1)
}

func (m *MockQuery) GetMany(ctx context.Context, arg repository.GetManyParams) ([]repository.User, error) {
	args := m.Called(ctx, arg)
	user, _ := args.Get(0).([]repository.User)
	return user, args.Error(1)
}

func (m *MockQuery) UpdateByID(ctx context.Context, arg repository.UpdateByIDParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}
