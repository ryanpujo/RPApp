package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello from another universe")
	})

	return route
}
