package registry

import "github.com/spriigan/broker/user/interface/controller"

func (r registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.client)
}
