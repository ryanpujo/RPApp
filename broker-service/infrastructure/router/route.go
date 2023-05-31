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

	productRoute := mux.Group("/api/product")
	productRoute.Use(gin.WrapH(ProductRoute(appController.Product, authentication.NewAuthentication())))

	userRoute := mux.Group("/api/user")
	userRoute.Use(gin.WrapH(UserRoute(appController.User, authentication.NewAuthentication())))

	return mux, func() {
		appController.Close()
	}
}
