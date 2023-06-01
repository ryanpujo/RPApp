package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/registry"
)

type Close func()

var mux = gin.Default()

func Route() (*gin.Engine, Close) {

	mux.GET("/api/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "helo")
	})
	register := registry.New()
	appController := register.NewAppController()
	auth := authentication.NewAuthentication()
	ProductRoute(appController.Product, auth)
	UserRoute(appController.User, auth)
	mux.POST("/api/register", auth.CreateUser)
	return mux, func() {
		appController.Close()
	}
}
