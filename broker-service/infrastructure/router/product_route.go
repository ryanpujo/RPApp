package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/interface/controller"
)

func ProductRoute(contr controller.ProductCrud, auth authentication.Authenticator) *gin.Engine {
	productRoute := mux.Group("/api/product")

	productRoute.Use(auth.Authenticate())
	productRoute.POST("/create", contr.Create)
	productRoute.GET("/:id", contr.GetById)
	productRoute.GET("/", contr.GetMany)
	productRoute.DELETE("/delete/:id", contr.DeleteById)
	productRoute.PATCH("/update", contr.UpdateById)
	return mux
}
