package main

import (
	"fmt"
	"log"

	"github.com/spriigan/RPApp/infrastructure"
	"github.com/spriigan/RPApp/registry"
)

func main() {
	app := infrastructure.Application()
	db := infrastructure.ConnectToDB()
	defer db.Close()
	register := registry.New(db)
	close, err := app.StartGrpcServer(register.NewUserServer())
	if err != nil {
		close()
		log.Fatal("failed to start the server", err)
	}
	fmt.Println("server started")
	defer close()
}
