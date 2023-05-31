package registry

import (
	"github.com/spriigan/broker/interface/controller"
	c "github.com/spriigan/broker/user/controller"
	"github.com/spriigan/broker/user/grpc/client"
)

func (r registry) NewUserClient() client.UserClientCloser {
	return client.NewUserClient()
}

func (r registry) NewUserController() controller.UserCrudCloser {
	return c.NewUserController(r.NewUserClient())
}
