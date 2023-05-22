package controller_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"
	"time"

	"github.com/ryanpujo/product-service/internal/controller"
	controllermocks "github.com/ryanpujo/product-service/internal/controller/controller_mocks"
	"github.com/ryanpujo/product-service/internal/pserror"
	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/ryanpujo/product-service/product-proto/grpc/product"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

var productServiceMock *controllermocks.ProductServiceMock
var lis *bufconn.Listener
var client product.ProductServiceClient

func buffDialer(ctx context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	lis = bufconn.Listen(1024 * 1024)
	defer lis.Close()
	s := grpc.NewServer()
	defer s.Stop()
	productServiceMock = new(controllermocks.ProductServiceMock)
	product.RegisterProductServiceServer(s, controller.NewProductController(productServiceMock))
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(buffDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client = product.NewProductServiceClient(conn)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
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
	createdProduct := repository.Product{
		StoreID:     1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryID:  1,
	}
	payload := product.ProductPayload{
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryId:  1,
		StoreId:     1,
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *product.CreatedProduct, err error)
	}{
		"product created": {
			arrange: func(t *testing.T) {
				productServiceMock.On("CreateProduct", mock.Anything, args).Return(createdProduct, nil).Once()
			},
			assert: func(t *testing.T, actual *product.CreatedProduct, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, createdProduct.Name, actual.Name)
			},
		},
		"failed to create": {
			arrange: func(t *testing.T) {
				productServiceMock.On("CreateProduct", mock.Anything, args).Return(nil, pserror.ErrNotFound).Once()
			},
			assert: func(t *testing.T, actual *product.CreatedProduct, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {

		t.Run(k, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
			defer cancel()
			v.arrange(t)

			result, err := client.CreateProduct(ctx, &payload)

			v.assert(t, result, err)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"deleted succesfully": {
			arrange: func(t *testing.T) {
				productServiceMock.On("DeleteProduct", mock.Anything, int64(1)).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"delete failed": {
			arrange: func(t *testing.T) {
				productServiceMock.On("DeleteProduct", mock.Anything, int64(1)).Return(pserror.ErrNotFound).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.DeleteProduct(context.Background(), &product.ProductID{Id: 1})

			v.assert(t, err)
		})
	}
}

func TestGetProductById(t *testing.T) {
	foundProduct := repository.GetProductByIdRow{
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
		assert  func(t *testing.T, actual *product.Product, err error)
	}{
		"found product": {
			arrange: func(t *testing.T) {
				productServiceMock.On("GetProductByID", mock.Anything, int64(1)).Return(foundProduct, nil).Once()
			},
			assert: func(t *testing.T, actual *product.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, foundProduct.Name, actual.Name)
			},
		},
		"product not found": {
			arrange: func(t *testing.T) {
				productServiceMock.On("GetProductByID", mock.Anything, int64(1)).Return(nil, pserror.ErrNotFound).Once()
			},
			assert: func(t *testing.T, actual *product.Product, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			actual, err := client.GetProductById(context.Background(), &product.ProductID{Id: 1})

			v.assert(t, actual, err)
		})
	}
}

func TestGetProducts(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *product.Products, err error)
	}{
		"succesfuly get products": {
			arrange: func(t *testing.T) {
				productServiceMock.On("GetProducts", mock.Anything, mock.Anything).Return([]repository.GetProductsRow{}, nil).Once()
			},
			assert: func(t *testing.T, actual *product.Products, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.IsType(t, &product.Products{}, actual)
			},
		},
		"failed to get products": {
			arrange: func(t *testing.T) {
				productServiceMock.On("GetProducts", mock.Anything, mock.Anything).Return(nil, errors.New(" an error")).Once()
			},
			assert: func(t *testing.T, actual *product.Products, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			products, err := client.GetProducts(context.Background(), &emptypb.Empty{})

			v.assert(t, products, err)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"product updated": {
			arrange: func(t *testing.T) {
				productServiceMock.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to update product": {
			arrange: func(t *testing.T) {
				productServiceMock.On("UpdateProduct", mock.Anything, mock.Anything).Return(pserror.ErrNotFound).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.UpdateProduct(context.Background(), &product.ProductPayload{})

			v.assert(t, err)
		})
	}
}
