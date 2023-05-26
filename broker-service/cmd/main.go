package main

import (
	"log"

	"github.com/spriigan/broker/infrastructure"
	"github.com/spriigan/broker/infrastructure/router"
)

func main() {
	app := infrastructure.Application()
	route, close := router.Route()
	defer close()
	if err := app.Serve(route); err != nil {
		log.Fatal(err)
	}
}
