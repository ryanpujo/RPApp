package mocks

import (
	"context"

	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockClient struct {
	mock.Mock
}

func (ps MockClient) CreateProduct(ctx context.Context, payload *product.ProductPayload, opts ...grpc.CallOption) (*product.CreatedProduct, error) {
	args := ps.Called(ctx, payload)
	created, _ := args.Get(0).(*product.CreatedProduct)
	return created, args.Error(1)
}

func (ps MockClient) GetProductById(ctx context.Context, id *product.ProductID, opts ...grpc.CallOption) (*product.Product, error) {
	args := ps.Called(ctx, id.Id)
	product, _ := args.Get(0).(*product.Product)
	return product, args.Error(1)
}

func (ps MockClient) GetProducts(ctx context.Context, empty *emptypb.Empty, opts ...grpc.CallOption) (*product.Products, error) {
	args := ps.Called(ctx, empty)
	products, _ := args.Get(0).(*product.Products)
	return products, args.Error(1)
}

func (ps MockClient) DeleteProduct(ctx context.Context, id *product.ProductID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := ps.Called(ctx, id.Id)
	empty, _ := args.Get(0).(*emptypb.Empty)
	return empty, args.Error(1)
}

func (ps MockClient) UpdateProduct(ctx context.Context, payload *product.ProductPayload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := ps.Called(ctx, payload)
	empty, _ := args.Get(0).(*emptypb.Empty)
	return empty, args.Error(1)
}

func (ps MockClient) Close() error {
	return nil
}
