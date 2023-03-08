package interactor

import (
	"context"
	"errors"

	"github.com/spriigan/RPApp/usecases/repository"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor interface {
	Create(ctx context.Context, user *models.UserPayload) (*models.UserBio, error)
	FindUsers(ctx context.Context) (*models.Users, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	DeleteByUsername(ctx context.Context, username string) error
	Update(ctx context.Context, user *models.UserPayload) error
}

var ErrDuplicateKeyInDatabase = errors.New("duplicate key in database")

type userInteractor struct {
	Repo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) *userInteractor {
	return &userInteractor{Repo: repo}
}

func (in *userInteractor) Create(ctx context.Context, user *models.UserPayload) (*models.UserBio, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	_, err := in.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user.GetBio(), nil
}

func (in *userInteractor) FindUsers(ctx context.Context) (*models.Users, error) {
	users, err := in.Repo.FindUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (in *userInteractor) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := in.Repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (in *userInteractor) DeleteByUsername(ctx context.Context, username string) error {
	err := in.Repo.DeleteByUsername(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

func (in *userInteractor) Update(ctx context.Context, user *models.UserPayload) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	err := in.Repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
