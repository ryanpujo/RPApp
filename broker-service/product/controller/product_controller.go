package controller

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	er "github.com/spriigan/broker/pkg/error"
	"github.com/spriigan/broker/product/domain"
	"github.com/spriigan/broker/product/grpc/client"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/spriigan/broker/response"
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

	result := domain.Product{
		ID:          createdProduct.Id,
		StoreID:     createdProduct.StoreId,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		ImageUrl:    createdProduct.ImageUrl,
		Stock:       int32(createdProduct.Stock),
		CategoryID:  createdProduct.CategoryId,
		CreatedAt:   sql.NullTime{Time: createdProduct.CreatedAt.AsTime()},
	}

	var res response.JsonRes
	res.Product = result

	c.JSON(http.StatusCreated, res)
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

	result := domain.Product{
		ID:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
		Stock:       int32(product.Stock),
		StoreName:   product.StoreName,
		Category:    product.Category,
		CreatedAt:   sql.NullTime{Time: product.CreatedAt.AsTime()},
	}

	var res response.JsonRes
	res.Product = result

	c.JSON(http.StatusOK, res)
}

func (p productController) GetMany(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	products, err := p.client.GetProducts(ctx, &emptypb.Empty{})
	if err != nil {
		er.Handle(c, err)
		return
	}

	var res response.JsonRes
	res.Products = []domain.Product{}

	if len(products.Products) == 0 {
		c.JSON(http.StatusOK, res)
		return
	}

	results := make([]domain.Product, 0, len(products.Products))

	for _, v := range products.Products {
		result := domain.Product{
			ID:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			ImageUrl:    v.ImageUrl,
			Stock:       int32(v.Stock),
			StoreName:   v.StoreName,
			Category:    v.Category,
			CreatedAt:   sql.NullTime{Time: v.CreatedAt.AsTime()},
		}

		results = append(results, result)
	}

	res.Products = results

	c.JSON(http.StatusOK, res)
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
