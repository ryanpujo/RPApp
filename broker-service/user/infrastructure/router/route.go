package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/user/interface/controller"
)

func Route(cont controller.UserController) *gin.Engine {
	mux := gin.Default()

	mux.POST("/", cont.Create)

	return mux
}
