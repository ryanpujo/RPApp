package service

import (
	"context"

	"github.com/ryanpujo/product-service/internal/pserror"
	"github.com/ryanpujo/product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, args repository.CreateProductParams) (repository.Product, error)
	GetProductByID(ctx context.Context, id int64) (repository.GetProductByIdRow, error)
	GetProducts(ctx context.Context) ([]repository.GetProductsRow, error)
	UpdateProduct(ctx context.Context, arg repository.UpdateProductParams) error
	DeleteProduct(ctx context.Context, id int64) error
}

type productService struct {
	query repository.QueriesInterface
}

func NewProductService(query repository.QueriesInterface) productService {
	return productService{query: query}
}

func (p productService) CreateProduct(ctx context.Context, args repository.CreateProductParams) (repository.Product, error) {
	createdProduct, err := p.query.CreateProduct(ctx, args)

	return createdProduct, pserror.ParseErrors(err)
}

func (p productService) DeleteProduct(ctx context.Context, id int64) error {
	_, err := p.query.GetProductById(ctx, id)
	if err != nil {
		return pserror.ParseErrors(err)
	}
	err = p.query.DeleteProduct(ctx, id)
	return pserror.ParseErrors(err)
}

func (p productService) GetProductByID(ctx context.Context, id int64) (repository.GetProductByIdRow, error) {
	product, err := p.query.GetProductById(ctx, id)
	return product, pserror.ParseErrors(err)
}

func (p productService) GetProducts(ctx context.Context) ([]repository.GetProductsRow, error) {
	products, err := p.query.GetProducts(ctx)
	return products, pserror.ParseErrors(err)
}

func (p productService) UpdateProduct(ctx context.Context, arg repository.UpdateProductParams) error {
	_, err := p.query.GetProductById(ctx, arg.ID)
	if err != nil {
		return pserror.ParseErrors(err)
	}
	err = p.query.UpdateProduct(ctx, arg)
	return pserror.ParseErrors(err)
}
