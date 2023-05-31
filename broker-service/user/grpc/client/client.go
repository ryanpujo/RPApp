package client

import (
	"fmt"
	"io"

	"github.com/spriigan/broker/infrastructure"
	"github.com/spriigan/broker/user/user-proto/userpb"
	"google.golang.org/grpc"
)

type UserClientCloser interface {
	userpb.UserServiceClient
	io.Closer
}

type userClient struct {
	userpb.UserServiceClient
	conn *grpc.ClientConn
}

func NewUserClient() *userClient {
	config := infrastructure.LoadConfig()
	service := config.Services["productservice"]
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", service.Address, service.ServicePort))
	if err != nil {
		panic(err)
	}
	return &userClient{UserServiceClient: userpb.NewUserServiceClient(conn), conn: conn}
}

func (u *userClient) Close() error {
	return u.conn.Close()
}
