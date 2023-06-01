package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/authentication/authmock"
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/product/controller"
	"github.com/spriigan/broker/product/controller/mocks"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/spriigan/broker/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	mockClient *mocks.MockClient
	mux        *gin.Engine
	mockAuth   *authmock.MockAuth
)

func TestMain(m *testing.M) {
	mockClient = new(mocks.MockClient)
	productController := controller.NewProductController(mockClient)
	defer productController.Close()
	mockAuth = new(authmock.MockAuth)
	auth := &authentication.Authentication{AuthClient: mockAuth}
	mux = router.ProductRoute(productController, auth)
	os.Exit(m.Run())
}

func unAuthorizeAssert(t *testing.T, statusCode int, json response.JsonRes) {
	require.Equal(t, http.StatusUnauthorized, statusCode)
	require.NotEmpty(t, json.Error)
	require.Equal(t, "unauthorized", json.Error)
}

func TestCreate(t *testing.T) {
	jsonReq := []byte(`{
		"store_id": 1,
		"name": "pujo",
		"description": "ryanpujo",
		"price": "4000",
		"image_url": "kjrkjnrjnrntkn",
		"stock": 9000,
		"category_id": 1
	}
	`)
	badJson := []byte(`{
		"store_id": 1,
		"name": "pujo",
		"description": "fdf",
		"price": "4000",
		"stock": 9000,
		"category_id": 1
	}
	`)
	testTable := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data response.JsonRes)
	}{
		"product created": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("CreateProduct", mock.Anything, mock.Anything).Return(&product.CreatedProduct{Name: "pujo"}, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, 201, statusCode)
				require.Empty(t, data.Error)
				require.Empty(t, data.Errors)
				require.NotEmpty(t, data.Product)
				require.Equal(t, "pujo", data.Product.Name)
			},
		},
		"vailidation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data)
				require.NotEmpty(t, data.Errors)
			},
		},
		"failed to create product": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("CreateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data)
				require.NotEmpty(t, data.Error)
				require.Equal(t, "an error", data.Error)
			},
		},
		"unauthorized acces": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("unauthorized")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewReader(v.json))
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer ksmdksmkdm")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestGetById(t *testing.T) {
	productTest := &product.Product{
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		Category:    "gadget",
		StoreName:   "ibox",
	}
	testTable := map[string]struct {
		id      int
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data response.JsonRes)
	}{
		"product found": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("GetProductById", mock.Anything, mock.Anything).Return(productTest, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.NotEmpty(t, data)
				require.Empty(t, data.Error)
				require.Empty(t, data.Errors)
				require.Equal(t, productTest.Name, data.Product.Name)
			},
		},
		"product not found": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("GetProductById", mock.Anything, mock.Anything).Return(nil, status.Error(codes.NotFound, "not found")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusNotFound, statusCode)
				require.Equal(t, "not found", data.Error)
			},
		},
		"bad uri": {
			id: 0,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data.Errors)
			},
		},
		"unauthorized": {
			id: 1,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("unauthorized")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%d", v.id), nil)
			req.Header.Add("Authorization", "Bearer ksmdksmkdm")
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestGetMany(t *testing.T) {
	productTest := &product.Product{
		Name:        "MacBook",
		Description: "good book",
		Price:       "3000",
		ImageUrl:    "dkskmf.com",
		Stock:       30,
		Category:    "gadget",
		StoreName:   "ibox",
	}
	products := &product.Products{
		Products: []*product.Product{productTest},
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data response.JsonRes)
	}{
		"products retrieved": {
			arrange: func(t *testing.T) {
				mockClient.On("GetProducts", mock.Anything, mock.Anything).Return(products, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data.Error)
				require.NotEmpty(t, data.Products)
				require.Equal(t, len(products.Products), len(data.Products))
			},
		},
		"failed to acces products": {
			arrange: func(t *testing.T) {
				mockClient.On("GetProducts", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.Empty(t, data.Products)
				require.NotEmpty(t, data.Error)
			},
		},
		"unauthorized": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodGet, "/", nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer kdfkdnfk")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestDeleteById(t *testing.T) {
	testTable := map[string]struct {
		id      int
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data response.JsonRes)
	}{
		"product deleted": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data.Error)
				require.Empty(t, data.Errors)
				require.Empty(t, data)
			},
		},
		"failed to delete": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data.Error)
			},
		},
		"validation error": {
			id: 0,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data.Errors)
			},
		},
		"unauthoriza": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/delete/%d", v.id), nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer jndjnjfdj")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestUpdateById(t *testing.T) {
	jsonReq := []byte(`{
		"id": 1,
		"store_id": 1,
		"name": "pujo",
		"description": "ryanpujo",
		"price": "4000",
		"image_url": "kjrkjnrjnrntkn",
		"stock": 9000,
		"category_id": 1
	}
	`)

	badJson := []byte(`{
		"id": 1,
		"store_id": 1,
		"name": "pujo",
		"price": "4000",
		"image_url": "kjrkjnrjnrntkn",
		"stock": 9000,
		"category_id": 1
	}
	`)
	testTable := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data response.JsonRes)
	}{
		"product updated": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data)
			},
		},
		"failed update": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data.Error)
				require.Equal(t, "an error", data.Error)
			},
		},
		"validation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data.Errors)
			},
		},
		"unauthorized": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodPatch, "/update", bytes.NewReader(v.json))
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer fnjkdnntr")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}
