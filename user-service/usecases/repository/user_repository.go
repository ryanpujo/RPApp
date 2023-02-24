package repository

import (
	"context"

	"github.com/spriigan/RPApp/user-proto/grpc/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserPayload) (int, error)
	FindUsers(ctx context.Context) (*models.Users, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	DeleteByUsername(ctx context.Context, username string) error
	Update(ctx context.Context, user *models.UserPayload) error
}
