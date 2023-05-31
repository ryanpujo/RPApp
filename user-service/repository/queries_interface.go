package repository

import "context"

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteByID(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (User, error)
	GetMany(ctx context.Context, args GetManyParams) ([]User, error)
	UpdateByID(ctx context.Context, arg UpdateByIDParams) error
}
