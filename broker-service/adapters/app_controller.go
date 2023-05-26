package adapters

import "github.com/spriigan/broker/interface/controller"

type AppController struct {
	Product controller.ProductCrud
}

func (app *AppController) Close() {
	app.Product.Close()
}
