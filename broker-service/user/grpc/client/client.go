package client

import (
	"github.com/spriigan/broker/user/user-proto/grpc/models"
	"google.golang.org/grpc"
)

type Close func()

func GrpcClient(addr string, opts ...grpc.DialOption) (models.UserServiceClient, Close, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, func() {
			conn.Close()
		}, err
	}
	return models.NewUserServiceClient(conn), func() {
		conn.Close()
	}, nil

}
