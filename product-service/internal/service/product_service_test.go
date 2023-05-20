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

func TestGetProductById(t *testing.T) {
	product := repository.GetProductByIdRow{
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual repository.GetProductByIdRow, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(product, nil).Once()
			},
			assert: func(t *testing.T, actual repository.GetProductByIdRow, err error) {
				require.NotEmpty(t, actual)
				require.NoError(t, err)
				require.Equal(t, product.Name, actual.Name)
			},
		},
		"failed call": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(nil, errors.New("product not found")).Once()
			},
			assert: func(t *testing.T, actual repository.GetProductByIdRow, err error) {
				require.Empty(t, actual)
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			product, err := productService.GetProductByID(context.Background(), int64(1))

			v.assert(t, product, err)
		})
	}
}

func TestGetProducts(t *testing.T) {
	product := repository.GetProductsRow{
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	products := []repository.GetProductsRow{product, product}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual []repository.GetProductsRow, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProducts", context.Background()).Return(products, nil).Once()
			},
			assert: func(t *testing.T, actual []repository.GetProductsRow, err error) {
				require.NotEmpty(t, actual)
				require.NoError(t, err)
				require.Equal(t, len(products), len(actual))
			},
		},
		"failed call": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProducts", context.Background()).Return(nil, errors.New("errors occur")).Once()
			},
			assert: func(t *testing.T, actual []repository.GetProductsRow, err error) {
				require.Empty(t, actual)
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			products, err := productService.GetProducts(context.Background())

			v.assert(t, products, err)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	product := repository.UpdateProductParams{
		ID:          1,
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"updated succefully": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(repository.GetProductByIdRow{}, nil).Once()
				queriesMock.On("UpdateProduct", context.Background(), product).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"no product found": {
			arrange: func(t *testing.T) {
				queriesMock.On("GetProductById", context.Background(), int64(1)).Return(repository.GetProductByIdRow{}, errors.New("product not found")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
				require.Equal(t, "product not found", err.Error())
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := productService.UpdateProduct(context.Background(), product)

			v.assert(t, err)
		})
	}
}
