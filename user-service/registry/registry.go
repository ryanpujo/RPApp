package registry

import (
	"database/sql"

	"github.com/spriigan/RPMedia/interface/controller"
	repo "github.com/spriigan/RPMedia/interface/repository"
	"github.com/spriigan/RPMedia/usecases/interactor"
	"github.com/spriigan/RPMedia/usecases/repository"
	"github.com/spriigan/RPMedia/user-proto/grpc/models"
)

type Registry interface {
	NewUserServer() models.UserServiceServer
}

type registry struct {
	DB *sql.DB
}

func New(db *sql.DB) *registry {
	return &registry{DB: db}
}

func (r *registry) NewUserServer() models.UserServiceServer {
	return controller.NewUserServer(r.newUserInteractor())
}

func (r *registry) newUserRepository() repository.UserRepository {
	return repo.NewUserRepository(r.DB)
}
func (r *registry) newUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.newUserRepository())
}
