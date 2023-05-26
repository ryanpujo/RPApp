package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/registry"
)

type Close func()

func Route() (*gin.Engine, Close) {
	mux := gin.Default()

	register := registry.New()
	appController := register.NewAppController()
	_ = mux.Group("/product", gin.WrapH(ProductRoute(appController.Product, authentication.NewAuthentication())))

	return mux, func() {
		appController.Close()
	}
}
