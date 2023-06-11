package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/interface/controller"
)

func UserRoute(contr controller.UserCrudCloser, auth authentication.Authenticator) *gin.Engine {
	userRoute := mux.Group("/api/user")

	userRoute.POST("/create", contr.Create)
	userRoute.GET("/:id", contr.GetById)
	userRoute.GET("/users", contr.GetMany)
	userRoute.DELETE("/delete/:id", contr.DeleteById)
	userRoute.PATCH("/update", contr.UpdateById)
	return mux
}
