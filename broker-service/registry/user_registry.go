package registry

import (
	"fmt"
	"github.com/spriigan/broker/infrastructure"
	"log"

	"github.com/spriigan/broker/user/grpc/client"
	"github.com/spriigan/broker/user/interface/controller"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (r registry) NewUserController() (controller.UserController, client.Close) {
	c, close := r.GrpcUserClient()
	return controller.NewUserController(c), close
}

func (r registry) GrpcUserClient() (models.UserServiceClient, client.Close) {
	config := infrastructure.LoadConfig()
	srv := config.Services["userservice"]
	fmt.Println(config)
	c, close, err := client.GrpcClient(fmt.Sprintf("%s:%d", srv.Address, srv.ServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return c, close
}
