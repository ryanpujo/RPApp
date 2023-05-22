package client_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/spriigan/broker/product/internal/grpc/client/mock/server"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	client     product.ProductServiceClient
	mockServer *server.ProductServerMock
	lis        *bufconn.Listener
)

func bufDialer(ctx context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	lis = bufconn.Listen(1024 * 1024)
	defer lis.Close()
	mockServer = new(server.ProductServerMock)
	s := grpc.NewServer()
	defer s.Stop()
	product.RegisterProductServiceServer(s, mockServer)
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	payload := &product.ProductPayload{
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		CategoryId:  1,
		StoreId:     1,
	}
	created := &product.CreatedProduct{
		Id:          1,
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
				mockServer.On("CreateProduct", mock.Anything, mock.Anything).Return(created, nil).Once()
			},
			assert: func(t *testing.T, actual *product.CreatedProduct, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, payload.Name, actual.Name)
			},
		},
		"failed to create": {
			arrange: func(t *testing.T) {
				mockServer.On("CreateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual *product.CreatedProduct, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			created, err := client.CreateProduct(context.Background(), payload)

			v.assert(t, created, err)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	foundProduct := &product.Product{
		Id:          1,
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		StoreName:   "IBox",
		Category:    "Gadget",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *product.Product, err error)
	}{
		"found product": {
			arrange: func(t *testing.T) {
				mockServer.On("GetProductById", mock.Anything, mock.Anything).Return(foundProduct, nil).Once()
			},
			assert: func(t *testing.T, actual *product.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, foundProduct.Name, actual.Name)
			},
		},
		"failed to got a product": {
			arrange: func(t *testing.T) {
				mockServer.On("GetProductById", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
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
	products := &product.Products{
		Products: []*product.Product{
			{
				Id:          1,
				Name:        "MacBook",
				Description: "good book",
				Price:       "3000",
				ImageUrl:    "dkskmf.com",
				Stock:       30,
				StoreName:   "IBox",
				Category:    "Gadget",
			},
		},
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *product.Products, err error)
	}{
		"success get products": {
			arrange: func(t *testing.T) {
				mockServer.On("GetProducts", mock.Anything, mock.Anything).Return(products, nil).Once()
			},
			assert: func(t *testing.T, actual *product.Products, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, len(products.Products), len(actual.Products))
			},
		},
		"failed to get products": {
			arrange: func(t *testing.T) {
				mockServer.On("GetProducts", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
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

func TestDeleteProduct(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"product deleted": {
			arrange: func(t *testing.T) {
				mockServer.On("DeleteProduct", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to delete": {
			arrange: func(t *testing.T) {
				mockServer.On("DeleteProduct", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
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

func TestUpdateProduct(t *testing.T) {
	payload := &product.ProductPayload{
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
		assert  func(t *testing.T, err error)
	}{
		"product updated": {
			arrange: func(t *testing.T) {
				mockServer.On("UpdateProduct", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to update": {
			arrange: func(t *testing.T) {
				mockServer.On("UpdateProduct", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.UpdateProduct(context.Background(), payload)

			v.assert(t, err)
		})
	}
}
