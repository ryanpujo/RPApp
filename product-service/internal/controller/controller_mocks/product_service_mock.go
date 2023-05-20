package controllermocks

import (
	"context"

	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (p ProductServiceMock) CreateProduct(ctx context.Context, args repository.CreateProductParams) (repository.Product, error) {
	arg := p.Called(ctx, args)
	product, _ := arg.Get(0).(repository.Product)
	return product, arg.Error(1)
}

func (p ProductServiceMock) DeleteProduct(ctx context.Context, id int64) error {
	args := p.Called(ctx, id)
	return args.Error(0)
}

func (p ProductServiceMock) GetProductByID(ctx context.Context, id int64) (repository.GetProductByIdRow, error) {
	args := p.Called(ctx, id)
	product, _ := args.Get(0).(repository.GetProductByIdRow)
	return product, args.Error(1)
}

func (p ProductServiceMock) GetProducts(ctx context.Context) ([]repository.GetProductsRow, error) {
	args := p.Called(ctx)
	products, _ := args.Get(0).([]repository.GetProductsRow)
	return products, args.Error(1)
}

func (p ProductServiceMock) UpdateProduct(ctx context.Context, arg repository.UpdateProductParams) error {
	args := p.Called(ctx, arg)
	return args.Error(0)
}
