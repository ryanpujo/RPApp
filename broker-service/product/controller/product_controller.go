package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	er "github.com/spriigan/broker/pkg/error"
	"github.com/spriigan/broker/product/domain"
	"github.com/spriigan/broker/product/grpc/client"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productController struct {
	client client.ProductServiceClientCloser
}

func NewProductController(client client.ProductServiceClientCloser) *productController {
	return &productController{client: client}
}

func (p productController) Create(c *gin.Context) {
	var json domain.Product
	err := c.ShouldBindJSON(&json)
	if err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	payload := product.ProductPayload{
		StoreId:     json.StoreID,
		Name:        json.Name,
		Description: json.Description,
		Price:       json.Price,
		ImageUrl:    json.ImageUrl,
		Stock:       uint32(json.Stock),
		CategoryId:  json.CategoryID,
	}
	createdProduct, err := p.client.CreateProduct(ctx, &payload)
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

type Uri struct {
	Id int64 `uri:"id" binding:"required,gt=0"`
}

func (p productController) GetById(c *gin.Context) {
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	product, err := p.client.GetProductById(ctx, &product.ProductID{Id: uri.Id})
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p productController) GetMany(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	products, err := p.client.GetProducts(ctx, &emptypb.Empty{})
	if err != nil {
		er.Handle(c, err)
		return
	}

	if len(products.Products) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []product.Product{}})
	}

	c.JSON(http.StatusOK, gin.H{"data": products.Products})
}

func (p productController) DeleteById(c *gin.Context) {
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	_, err := p.client.DeleteProduct(ctx, &product.ProductID{Id: uri.Id})
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (p productController) UpdateById(c *gin.Context) {
	var json domain.Product
	if err := c.ShouldBindJSON(&json); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	payload := &product.ProductPayload{
		Id:          json.ID,
		StoreId:     json.StoreID,
		Name:        json.Name,
		Description: json.Description,
		Price:       json.Price,
		ImageUrl:    json.ImageUrl,
		Stock:       uint32(json.Stock),
		CategoryId:  json.CategoryID,
	}

	_, err := p.client.UpdateProduct(ctx, payload)
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (p *productController) Close() error {
	return p.client.Close()
}
