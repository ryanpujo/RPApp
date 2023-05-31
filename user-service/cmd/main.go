package main

import (
	"github.com/spriigan/RPApp/infrastructure"
	"github.com/spriigan/RPApp/registry"
)

func main() {
	app := infrastructure.Application()
	db := app.ConnectToDB()
	defer db.Close()
	r := registry.New(db)
	close, err := app.StartGrpcServer(r.NewUserController())
	if err != nil {
		close()
		panic(err)
	}
	defer close()
}
