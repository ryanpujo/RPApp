package registry

import (
	"github.com/spriigan/broker/adapters"
)

type Registry interface {
	NewAppController() *adapters.AppController
}

type registry struct {
}

func New() *registry {
	return &registry{}
}

func (r registry) NewAppController() *adapters.AppController {
	return &adapters.AppController{
		Product: r.NewProductController(),
		User:    r.NewUserController(),
	}
}
