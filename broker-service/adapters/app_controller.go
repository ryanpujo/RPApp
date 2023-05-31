package adapters

import "github.com/spriigan/broker/interface/controller"

type AppController struct {
	Product controller.ProductCrud
	User    controller.UserCrudCloser
}

func (app *AppController) Close() {
	app.User.Close()
	app.Product.Close()
}
