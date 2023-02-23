package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/adapters"
)

func Route(cont *adapters.AppController) *gin.Engine {
	mux := gin.Default()

	mux.POST("/user", cont.User.Create)

	return mux
}
