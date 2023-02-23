package registry

import (
	"github.com/spriigan/broker/adapters"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
)

type Registry interface {
	NewAppController() *adapters.AppController
}

type registry struct {
	client models.UserServiceClient
}

func New(c models.UserServiceClient) *registry {
	return &registry{client: c}
}

func (r registry) NewAppController() *adapters.AppController {
	return &adapters.AppController{User: r.NewUserController()}
}
