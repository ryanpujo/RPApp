package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/interface/controller"
)

func UserRoute(contr controller.UserCrudCloser, auth authentication.Authenticator) *gin.Engine {
	mux := gin.Default()

	mux.Use(auth.Authenticate())
	mux.POST("/create", contr.Create)
	mux.GET("/:id", contr.GetById)
	mux.GET("/users/:limit", contr.GetMany)
	mux.DELETE("/delete/:id", contr.DeleteById)
	mux.PATCH("/update", contr.UpdateById)
	return mux
}