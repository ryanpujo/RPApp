package main

import (
	"log"

	"github.com/spriigan/broker/infrastructure"
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/registry"
)

func main() {
	app := infrastructure.Application()
	register := registry.New()
	appController, close := register.NewAppController()
	defer close()
	if err := app.Serve(router.Route(appController)); err != nil {
		log.Fatal(err)
	}
}
