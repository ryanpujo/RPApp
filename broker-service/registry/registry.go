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

func (r registry) NewAppController() *adapters.AppController {
	return &adapters.AppController{Product: r.NewProductController()}
}
