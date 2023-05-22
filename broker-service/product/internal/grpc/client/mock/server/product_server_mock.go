package server

import (
	"context"

	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductServerMock struct {
	mock.Mock
	product.UnimplementedProductServiceServer
}

func (ps ProductServerMock) CreateProduct(ctx context.Context, payload *product.ProductPayload) (*product.CreatedProduct, error) {
	args := ps.Called(ctx, payload)
	created, _ := args.Get(0).(*product.CreatedProduct)
	return created, args.Error(1)
}

func (ps ProductServerMock) GetProductById(ctx context.Context, id *product.ProductID) (*product.Product, error) {
	args := ps.Called(ctx, id.Id)
	product, _ := args.Get(0).(*product.Product)
	return product, args.Error(1)
}

func (ps ProductServerMock) GetProducts(ctx context.Context, empty *emptypb.Empty) (*product.Products, error) {
	args := ps.Called(ctx, empty)
	products, _ := args.Get(0).(*product.Products)
	return products, args.Error(1)
}

func (ps ProductServerMock) DeleteProduct(ctx context.Context, id *product.ProductID) (*emptypb.Empty, error) {
	args := ps.Called(ctx, id.Id)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}

func (ps ProductServerMock) UpdateProduct(ctx context.Context, payload *product.ProductPayload) (*emptypb.Empty, error) {
	args := ps.Called(ctx, payload)
	return args.Get(0).(*emptypb.Empty), args.Error(1)
}
