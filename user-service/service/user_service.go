package service

import (
	"context"

	"github.com/spriigan/RPApp/pkg/usererror"
	"github.com/spriigan/RPApp/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error)
	DeleteByID(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (repository.User, error)
	GetMany(ctx context.Context, limit int32) ([]repository.User, error)
	UpdateByID(ctx context.Context, arg repository.UpdateByIDParams) error
}

type userService struct {
	query repository.Querier
}

func NewUserService(query repository.Querier) *userService {
	return &userService{query: query}
}

func (u userService) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	createdUser, err := u.query.CreateUser(ctx, arg)
	return createdUser, usererror.ParseErrors(err)
}

func (u userService) DeleteByID(ctx context.Context, id int64) error {
	err := u.query.DeleteByID(ctx, id)
	return usererror.ParseErrors(err)
}

func (u userService) GetById(ctx context.Context, id int64) (repository.User, error) {
	user, err := u.query.GetById(ctx, id)
	return user, usererror.ParseErrors(err)
}

func (u userService) GetMany(ctx context.Context, limit int32) ([]repository.User, error) {
	users, err := u.query.GetMany(ctx, limit)
	return users, usererror.ParseErrors(err)
}

func (u userService) UpdateByID(ctx context.Context, arg repository.UpdateByIDParams) error {
	err := u.query.UpdateByID(ctx, arg)
	return usererror.ParseErrors(err)
}
