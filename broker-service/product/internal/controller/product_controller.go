package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

type productController struct {
	client product.ProductServiceClient
}
