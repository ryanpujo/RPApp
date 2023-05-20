package infra

import (
	"fmt"
	"net"

	"github.com/ryanpujo/product-service/product-proto/grpc/product"
	"google.golang.org/grpc"
)

type application struct {
	config Config
}

func Application() application {
	return application{config: LoadConfig()}
}

type close func()

func (app application) StartGrpcServer(server product.ProductServiceServer) (close, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.config.Port))
	if err != nil {
		return func() { lis.Close() }, err
	}

	s := grpc.NewServer()
	product.RegisterProductServiceServer(s, server)

	if err := s.Serve(lis); err != nil {
		return func() {
			lis.Close()
			s.Stop()
		}, err
	}

	return func() {
		lis.Close()
		s.Stop()
	}, err
}
