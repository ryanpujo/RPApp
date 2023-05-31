package mocks

import (
	"context"

	"github.com/spriigan/broker/user/user-proto/userpb"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockUserServer struct {
	userpb.UnimplementedUserServiceServer
	mock.Mock
}

func (m MockUserServer) CreateUser(ctx context.Context, in *userpb.UserPayload) (*userpb.User, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.User)
	return user, args.Error(1)
}

func (m MockUserServer) GetMany(ctx context.Context, in *userpb.Limit) (*userpb.Users, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.Users)
	return user, args.Error(1)
}

func (m MockUserServer) GetById(ctx context.Context, in *userpb.UserId) (*userpb.User, error) {
	args := m.Called(ctx, in)
	user, _ := args.Get(0).(*userpb.User)
	return user, args.Error(1)
}

func (m MockUserServer) DeleteById(ctx context.Context, in *userpb.UserId) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (m MockUserServer) UpdateById(ctx context.Context, in *userpb.UserPayload) (*emptypb.Empty, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}
