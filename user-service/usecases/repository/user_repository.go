package repository

import (
	"github.com/spriigan/RPApp/domain"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
)

type UserRepository interface {
	Create(user *models.UserPayload) (int, error)
	FindUsers() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	DeleteByUsername(username string) error
	Update(user domain.UserPayload) error
}
