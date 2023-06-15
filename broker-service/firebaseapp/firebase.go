package firebaseapp

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func New() *firebase.App {
	if app != nil {
		return app
	}
	config := firebase.Config{
		ProjectID:     "orbit-app-145b9",
		StorageBucket: "orbit-app-145b9.appspot.com",
	}
	opt := option.WithCredentialsFile("./orbit-app-145b9-firebase-adminsdk-7ycvp-6ab97f8272.json")
	app, err := firebase.NewApp(context.Background(), &config, opt)
	if err != nil {
		log.Fatal(err)
	}
	return app
}
