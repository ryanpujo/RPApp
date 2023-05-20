package mocks

import (
	"context"
	"database/sql"

	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/stretchr/testify/mock"
)

type QueriesMock struct {
	mock.Mock
}

func (qm *QueriesMock) CreateProduct(ctx context.Context, args repository.CreateProductParams) (repository.Product, error) {
	returned := qm.Called(ctx, args)
	product, ok := returned.Get(0).(repository.Product)
	if ok {
		return product, returned.Error(1)
	}
	return product, returned.Error(1)
}

func (qm *QueriesMock) GetProductById(ctx context.Context, id int64) (repository.GetProductByIdRow, error) {
	returned := qm.Called(ctx, id)
	product, _ := returned.Get(0).(repository.GetProductByIdRow)
	return product, returned.Error(1)
}

func (qm *QueriesMock) GetProducts(ctx context.Context) ([]repository.GetProductsRow, error) {
	returned := qm.Called(ctx)
	products, _ := returned.Get(0).([]repository.GetProductsRow)
	return products, returned.Error(1)
}

func (qm *QueriesMock) UpdateProduct(ctx context.Context, arg repository.UpdateProductParams) error {
	returned := qm.Called(ctx, arg)
	return returned.Error(0)
}

func (qm *QueriesMock) DeleteProduct(ctx context.Context, id int64) error {
	returned := qm.Called(ctx, id)
	return returned.Error(0)
}

func (qm *QueriesMock) WithTx(tx *sql.Tx) *repository.Queries {
	return &repository.Queries{}
}
