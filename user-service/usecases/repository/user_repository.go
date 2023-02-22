package repository

import "github.com/spriigan/RPMedia/domain"

type UserRepository interface {
	Create(user *domain.UserPayload) (int, error)
	FindUsers() ([]*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	DeleteByUsername(username string) error
	Update(user domain.UserPayload) error
}
