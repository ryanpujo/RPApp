package main

import (
	"log"

	"github.com/spriigan/broker/infrastructure"
	"github.com/spriigan/broker/infrastructure/grpc/client"
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := infrastructure.Application()
	client, close, err := client.GrpcClient("user-service:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		close()
		log.Fatal(err)
	}
	defer close()
	register := registry.New(client)
	if err := app.Serve(router.Route(register.NewAppController())); err != nil {
		log.Fatal(err)
	}
}
