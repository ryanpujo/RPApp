package registry

import (
	"github.com/spriigan/broker/adapters"
	"github.com/spriigan/broker/user/grpc/client"
)

type Registry interface {
	NewAppController() (*adapters.AppController, client.Close)
}

type registry struct {
}

func New() *registry {
	return &registry{}
}

func (r registry) NewAppController() (*adapters.AppController, client.Close) {
	user, close := r.NewUserController()
	return &adapters.AppController{User: user}, func() {
		close()
	}
}
