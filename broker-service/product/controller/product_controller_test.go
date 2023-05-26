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
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/product/controller"
	"github.com/spriigan/broker/product/controller/mocks"
	"github.com/spriigan/broker/product/product-proto/grpc/product"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	mockClient *mocks.MockClient
	mux        *gin.Engine
	mockAuth   *mocks.MockAuth
)

func TestMain(m *testing.M) {
	mockClient = new(mocks.MockClient)
	productController := controller.NewProductController(mockClient)
	defer productController.Close()
	mockAuth = new(mocks.MockAuth)
	auth := &authentication.Authentication{AuthClient: mockAuth}
	mux = router.ProductRoute(productController, auth)
	os.Exit(m.Run())
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
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"product created": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("CreateProduct", mock.Anything, mock.Anything).Return(&product.CreatedProduct{Name: "pujo"}, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, 201, statusCode)
				require.Empty(t, data["error"])
				require.Empty(t, data["errors"])
				require.NotEmpty(t, data)
				require.Equal(t, "pujo", data["name"])
			},
		},
		"vailidation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data)
				require.NotEmpty(t, data["errors"])
			},
		},
		"failed to create product": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("CreateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "an error", data["error"])
			},
		},
		"unauthorized acces": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("unauthorized")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "unathorized", data["error"])
			},
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
			var res gin.H
			_ = json.NewDecoder(rr.Body).Decode(&res)
			print(res)

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
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"product found": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("GetProductById", mock.Anything, mock.Anything).Return(productTest, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusOK, statusCode)
				require.NotEmpty(t, data)
				require.Empty(t, data["error"])
				require.Empty(t, data["errors"])
				require.Equal(t, productTest.Name, data["name"])
			},
		},
		"product not found": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("GetProductById", mock.Anything, mock.Anything).Return(nil, status.Error(codes.NotFound, "not found")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusNotFound, statusCode)
				require.Equal(t, "not found", data["error"])
			},
		},
		"bad uri": {
			id: 0,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data["errors"])
			},
		},
		"unauthorized": {
			id: 1,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("unauthorized")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "unathorized", data["error"])
			},
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
			var res gin.H
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
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"products retrieved": {
			arrange: func(t *testing.T) {
				mockClient.On("GetProducts", mock.Anything, mock.Anything).Return(products, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data["error"])
				require.NotEmpty(t, data["data"])
				require.Equal(t, len(products.Products), len(data["data"].([]interface{})))
			},
		},
		"failed to acces products": {
			arrange: func(t *testing.T) {
				mockClient.On("GetProducts", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.Empty(t, data["data"])
				require.NotEmpty(t, data["error"])
			},
		},
		"unauthorized": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "unathorized", data["error"])
			},
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
			var res gin.H
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestDeleteById(t *testing.T) {
	testTable := map[string]struct {
		id      int
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"product deleted": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data["error"])
				require.Empty(t, data["errors"])
				require.Empty(t, data)
			},
		},
		"failed to delete": {
			id: 1,
			arrange: func(t *testing.T) {
				mockClient.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data["error"])
			},
		},
		"validation error": {
			id: 0,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data["errors"])
			},
		},
		"unauthoriza": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "unathorized", data["error"])
			},
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
			var res gin.H
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
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"product updated": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil, nil).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, data)
				require.Empty(t, data["error"])
				require.Empty(t, data["errors"])
			},
		},
		"failed update": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				mockClient.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "an error", data["error"])
			},
		},
		"validation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(&auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, data["errors"])
			},
		},
		"unauthorized": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusUnauthorized, statusCode)
				require.NotEmpty(t, data["error"])
				require.Equal(t, "unathorized", data["error"])
			},
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
			var res gin.H
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}
