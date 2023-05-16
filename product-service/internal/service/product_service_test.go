package service_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/ryanpujo/product-service/internal/service"
	"github.com/ryanpujo/product-service/internal/service/mocks"
	"github.com/stretchr/testify/require"
)

var productService service.ProductService
var queriesMock *mocks.QueriesMock

func TestMain(m *testing.M) {
	queriesMock = new(mocks.QueriesMock)
	productService = service.NewProductService(queriesMock)
	os.Exit(m.Run())
}

func TestCreateProduct(t *testing.T) {
	args := repository.CreateProductParams{
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	product := repository.Product{
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	pqErr := &pgconn.PgError{Code: "23502", ColumnName: "name", Message: "not nul violation"}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual repository.Product, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				queriesMock.On("CreateProduct", context.Background(), args).Return(product, nil).Once()
			},
			assert: func(t *testing.T, actual repository.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, args.Name, actual.Name)
			},
		},
		"failed test": {
			arrange: func(t *testing.T) {
				queriesMock.On("CreateProduct", context.Background(), args).Return(nil, pqErr).Once()
			},
			assert: func(t *testing.T, actual repository.Product, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
				require.Equal(t, fmt.Sprintf("make sure all required field is filled:%s", pqErr.Error()), err.Error())
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			product, err := productService.CreateProduct(context.Background(), args)

			v.assert(t, product, err)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"deleted successfully": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(repository.GetProductByIdRow{}, nil).Once()
				queriesMock.On("DeleteProduct", context.Background(), int64(1)).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"product not found": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(nil, errors.New("product not found")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
				require.Equal(t, "product not found", err.Error())
			},
		},
		"delate failed": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(repository.GetProductByIdRow{}, nil).Once()
				queriesMock.On("DeleteProduct", context.Background(), int64(1)).Return(errors.New("failed to delete")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
				require.Equal(t, "failed to delete", err.Error())
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := productService.DeleteProduct(context.Background(), 1)

			v.assert(t, err)
		})
	}
}
