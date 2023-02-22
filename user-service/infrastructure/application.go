package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spriigan/RPMedia/user-proto/grpc/models"
	"google.golang.org/grpc"
)

type application struct {
	Congfig congfig
}

func Application() application {
	return application{
		Congfig: congfig{
			GRPC_PORT: os.Getenv("GRPC_PORT"),
			DSN:       os.Getenv("DSN"),
		},
	}
}

func (app *application) StartGrpcServer(server models.UserServiceServer) (func(), error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.Congfig.GRPC_PORT))
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) ConnectToDB() *sql.DB {
	ticker := time.NewTicker(1 * time.Second)
	var db *sql.DB
	var err error
	count := 0

	for db == nil {
		db, err = openDB(app.Congfig.DSN)
		if err != nil {
			log.Println("postgres is not ready yet:", err)
		}
		count++
		if count > 5 {
			log.Fatal("cant connect to postgres:", err)
		}
		<-ticker.C
	}
	return db
}
