package client

import (
	"fmt"
	"io"

	"github.com/spriigan/broker/infrastructure"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"google.golang.org/grpc"
)

type ProductServiceClientCloser interface {
	product.ProductServiceClient
	io.Closer
}

type productClient struct {
	product.ProductServiceClient
	conn *grpc.ClientConn
}

func NewProductClient() *productClient {
	config := infrastructure.LoadConfig()
	service := config.Services["productservice"]
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.ServicePort))
	if err != nil {
		panic(err)
	}
	return &productClient{ProductServiceClient: product.NewProductServiceClient(conn), conn: conn}
}

func (pc *productClient) Close() error {
	return pc.conn.Close()
}
