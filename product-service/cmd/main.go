package main

import (
	"github.com/ryanpujo/product-service/internal/controller"
	"github.com/ryanpujo/product-service/internal/infra"
	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/ryanpujo/product-service/internal/service"
)

func main() {
	app := infra.Application()
	db := app.ConnectToDB()
	defer db.Close()
	repo := repository.New(db)
	productService := service.NewProductService(repo)
	productController := controller.NewProductController(productService)
	close, err := app.StartGrpcServer(productController)
	if err != nil {
		close()
		panic(err)
	}
	defer close()
}
