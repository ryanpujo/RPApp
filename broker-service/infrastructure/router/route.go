package router

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/adapters"
	"github.com/spriigan/broker/authentication"
	"google.golang.org/api/option"
)

func Route(cont *adapters.AppController) *gin.Engine {
	mux := gin.Default()

	config := firebase.Config{
		ProjectID: "orbit-app-145b9",
	}
	opt := option.WithCredentialsFile("./orbit-app-145b9-firebase-adminsdk-7ycvp-6ab97f8272.json")
	app, err := firebase.NewApp(context.Background(), &config, opt)
	if err != nil {
		log.Fatal(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := authentication.NewAuthentication(authClient)

	protected := mux.Group("/auth")
	protected.Use(auth.Authenticate())
	{
		protected.GET("/user", cont.User.FindUsers)
		protected.GET("/user/:username", cont.User.FindByUsername)
		protected.DELETE("/user/:username", cont.User.DeleteByUsername)
		protected.PATCH("/user", cont.User.Update)
	}
	public := mux.Group("/public")
	public.POST("/user", cont.User.Create)
	return mux
}
