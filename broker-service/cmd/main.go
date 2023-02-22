package main

import (
	"log"

	"github.com/spriigan/broker/user/infrastructure"
	"github.com/spriigan/broker/user/infrastructure/grpc/client"
	"github.com/spriigan/broker/user/infrastructure/router"
	"github.com/spriigan/broker/user/interface/controller"
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
	if err := app.Serve(router.Route(controller.NewUserController(client))); err != nil {
		log.Fatal(err)
	}
}
