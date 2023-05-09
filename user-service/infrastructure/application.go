package infrastructure

import (
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"google.golang.org/grpc"
	"net"
)

type application struct {
	Config config
}

func Application() application {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	return application{
		Config: config{
			GRPC_PORT: viper.GetInt("port"),
		},
	}
}

func (app *application) StartGrpcServer(server models.UserServiceServer) (func(), error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.GRPC_PORT))
	if err != nil {
		return func() {
			lis.Close()
		}, err
	}
	s := grpc.NewServer()
	models.RegisterUserServiceServer(s, server)

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
