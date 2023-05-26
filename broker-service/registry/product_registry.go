package registry

import (
	c "github.com/spriigan/broker/interface/controller"
	"github.com/spriigan/broker/product/controller"
	"github.com/spriigan/broker/product/grpc/client"
)

func (r registry) NewProductClient() client.ProductServiceClientCloser {
	return client.NewProductClient()
}

func (r registry) NewProductController() c.ProductCrud {
	return controller.NewProductController(r.NewProductClient())
}
