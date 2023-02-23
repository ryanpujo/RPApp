package interactor

import (
	"github.com/spriigan/RPApp/usecases/repository"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor interface {
	Create(user *models.UserPayload) (int, error)
	FindUsers() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	DeleteByUsername(username string) error
	Update(user *models.UserPayload) error
}

type userInteractor struct {
	Repo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) *userInteractor {
	return &userInteractor{Repo: repo}
}

func (in *userInteractor) Create(user *models.UserPayload) (int, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	id, err := in.Repo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (in *userInteractor) FindUsers() (*models.Users, error) {
	users, err := in.Repo.FindUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (in *userInteractor) FindByUsername(username string) (*models.User, error) {
	user, err := in.Repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (in *userInteractor) DeleteByUsername(username string) error {
	err := in.Repo.DeleteByUsername(username)
	if err != nil {
		return err
	}
	return nil
}

func (in *userInteractor) Update(user *models.UserPayload) error {
	err := in.Repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
