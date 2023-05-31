package registry

import (
	"database/sql"

	"github.com/spriigan/RPApp/controller"
	"github.com/spriigan/RPApp/repository"
	"github.com/spriigan/RPApp/service"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"github.com/spriigan/RPApp/user-proto/userpb"
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

func (r registry) NewUserRepository() repository.Querier {
	return repository.New(r.DB)
}

func (r registry) NewUserService() service.UserService {
	return service.NewUserService(r.NewUserRepository())
}

func (r registry) NewUserController() userpb.UserServiceServer {
	return controller.NewUserController(r.NewUserService())
}
