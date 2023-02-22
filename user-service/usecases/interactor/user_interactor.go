package interactor

import (
	"github.com/spriigan/RPMedia/domain"
	"github.com/spriigan/RPMedia/usecases/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor interface {
	Create(user *domain.UserPayload) (int, error)
	FindUsers() ([]*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	DeleteByUsername(username string) error
	Update(user domain.UserPayload) error
}

type userInteractor struct {
	Repo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) *userInteractor {
	return &userInteractor{Repo: repo}
}

func (in *userInteractor) Create(user *domain.UserPayload) (int, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	id, err := in.Repo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (in *userInteractor) FindUsers() ([]*domain.User, error) {
	users, err := in.Repo.FindUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (in *userInteractor) FindByUsername(username string) (*domain.User, error) {
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

func (in *userInteractor) Update(user domain.UserPayload) error {
	err := in.Repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
