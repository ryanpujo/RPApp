package adapters

import "github.com/spriigan/broker/user/interface/controller"

type AppController struct {
	User interface{ controller.UserController }
}
