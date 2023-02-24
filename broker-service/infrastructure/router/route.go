package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/adapters"
)

func Route(cont *adapters.AppController) *gin.Engine {
	mux := gin.Default()

	mux.POST("/user", cont.User.Create)
	mux.GET("/user", cont.User.FindUsers)
	mux.GET("/user/:username", cont.User.FindByUsername)
	mux.DELETE("/user/:username", cont.User.DeleteByUsername)
	mux.PATCH("/user", cont.User.Update)

	return mux
}
