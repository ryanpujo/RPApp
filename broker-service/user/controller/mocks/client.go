package mocks

import (
	"context"

	"github.com/spriigan/broker/user/user-proto/userpb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) CreateUser(ctx context.Context, in *userpb.UserPayload, opts ...grpc.CallOption) (*userpb.User, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.User)
	return user, args.Error(1)
}

func (m *MockClient) GetMany(ctx context.Context, in *userpb.Limit, opts ...grpc.CallOption) (*userpb.Users, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.Users)
	return user, args.Error(1)
}

func (m *MockClient) GetById(ctx context.Context, in *userpb.UserId, opts ...grpc.CallOption) (*userpb.User, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.User)
	return user, args.Error(1)
}

func (m *MockClient) DeleteById(ctx context.Context, in *userpb.UserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockClient) UpdateById(ctx context.Context, in *userpb.UserPayload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m *MockClient) Close() error {
	return nil
}
