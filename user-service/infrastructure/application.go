package infrastructure

import (
	"fmt"
	"net"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spriigan/RPApp/user-proto/userpb"
	"google.golang.org/grpc"
)

type application struct {
	Config config
}

func Application() application {
	LoadConfig()
	return application{}
}

func (app *application) StartGrpcServer(server userpb.UserServiceServer) (func(), error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", Cfg.PORT))
	if err != nil {
		return func() {
			lis.Close()
		}, err
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, server)

	if err = s.Serve(lis); err != nil {
		return func() {
			lis.Close()
			s.Stop()
		}, err
	}

	return func() {
		lis.Close()
		s.Stop()
	}, nil
}
