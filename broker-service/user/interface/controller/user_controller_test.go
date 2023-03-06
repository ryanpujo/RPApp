package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/adapters"
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/interface/controller"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type mockClient struct {
	mock.Mock
}

func (mc mockClient) RegisterUser(ctx context.Context, in *models.UserPayload, opts ...grpc.CallOption) (*models.UserBio, error) {
	args := mc.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserBio), args.Error(1)
}

func (mc mockClient) FindUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*models.Users, error) {
	args := mc.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Users), args.Error(1)
}

func (mc mockClient) FindByUsername(ctx context.Context, in *models.Username, opts ...grpc.CallOption) (*models.UserBio, error) {
	args := mc.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserBio), args.Error(1)
}

func (mc mockClient) DeleteByUsername(ctx context.Context, in *models.Username, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := mc.Called(ctx, in)
	return nil, args.Error(1)
}

func (mc mockClient) Update(ctx context.Context, in *models.UserPayload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	args := mc.Called(ctx, in)
	return nil, args.Error(1)
}

var ac *adapters.AppController
var client *mockClient
var mux *gin.Engine

func TestMain(m *testing.M) {
	client = new(mockClient)
	ac = &adapters.AppController{
		User: controller.NewUserController(client),
	}
	mux = router.Route(ac)
	os.Exit(m.Run())
}

func TestRegisterUser(t *testing.T) {
	jsonReq := []byte(`{
		"fname": "ryan",
		"lname": "pujo",
		"username": "ryanpujo",
		"email": "ryanpuj@ogmail.com",
		"password": "kjrkjnrjnrntkn"
	}
	`)

	wrongValidation := []byte(`{
		"fname": "ryan",
		"lname": "pujo",
		"username": "ryanpujo",
		"email": "ryanpuj@ogmail.com",
		"password": "fdf"
	}
	`)
	testTabel := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data gin.H)
	}{
		"success api call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("RegisterUser", mock.Anything, mock.Anything).Return(&models.UserBio{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusCreated, statusCode)
				require.NotZero(t, data["data"])
			},
		},
		"failed call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("RegisterUser", mock.Anything, mock.Anything).Return(nil, status.Error(codes.FailedPrecondition, "got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Zero(t, data["data"])
			},
		},
		"wrong validation": {
			json:    wrongValidation,
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, statusCode int, data gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Zero(t, data["data"])
			},
		},
	}

	for k, v := range testTabel {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(v.json))
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res gin.H
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestFindUsers(t *testing.T) {
	users := &models.Users{
		User: []*models.UserBio{
			{},
			{},
			{},
		},
	}
	testTabel := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data interface{}, isError bool)
	}{
		"success api call": {
			arrange: func(t *testing.T) {
				client.On("FindUsers", mock.Anything, mock.Anything).Return(users, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusOK, statusCode)
				require.False(t, isError)
				require.NotNil(t, data)
			},
		},
		"failed call": {
			arrange: func(t *testing.T) {
				client.On("FindUsers", mock.Anything, mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Nil(t, data)
				require.True(t, isError)
			},
		},
	}

	for k, v := range testTabel {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, _ := http.NewRequest(http.MethodGet, "/user", nil)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonResponse
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res.Data, res.Error)
		})
	}
}

func TestFindByUsername(t *testing.T) {
	user := &models.UserBio{Lname: "connor"}
	testTabel := map[string]struct {
		uri     string
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, data interface{}, isError bool)
	}{
		"success api call": {
			uri: "/user/ryanpuj0",
			arrange: func(t *testing.T) {
				client.On("FindByUsername", mock.Anything, mock.Anything).Return(user, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusOK, statusCode)
				require.False(t, isError)
				require.NotNil(t, data)
			},
		},
		"failed call": {
			uri: "/user/ryanpujo",
			arrange: func(t *testing.T) {
				client.On("FindByUsername", mock.Anything, mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Zero(t, data)
				require.True(t, isError)
			},
		},
		"bad uri": {
			uri:     "/user/rt",
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Nil(t, data)
				require.True(t, isError)
			},
		},
	}

	for k, v := range testTabel {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, _ := http.NewRequest(http.MethodGet, v.uri, nil)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonResponse
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res.Data, res.Error)
		})
	}
}

func TestDeleteByUsername(t *testing.T) {
	testTable := map[string]struct {
		uri     string
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, message string, isError bool)
	}{
		"success api call": {
			uri: "/user/ryanpuj0",
			arrange: func(t *testing.T) {
				client.On("DeleteByUsername", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusOK, statusCode)
				require.False(t, isError)
				require.NotNil(t, message)
				require.Equal(t, "user has been deleted", message)
			},
		},
		"failed call": {
			uri: "/user/ryanpujo",
			arrange: func(t *testing.T) {
				client.On("DeleteByUsername", mock.Anything, mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, message)
				require.Equal(t, "got an error", message)
				require.True(t, isError)
			},
		},
		"bad uri": {
			uri:     "/user/rt",
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, message)
				require.True(t, isError)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, _ := http.NewRequest(http.MethodDelete, v.uri, nil)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonResponse
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res.Message, res.Error)
		})
	}
}

func TestUpdate(t *testing.T) {
	jsonReq := []byte(`{
		"fname": "ryan",
		"lname": "pujo",
		"username": "ryanpujo",
		"email": "ryanpuj@ogmail.com",
		"password": "kjrkjnrjnrntkn"
	}
	`)

	wrongValidation := []byte(`{
		"fname": "ryan",
		"lname": "pujo",
		"username": "ryanpujo",
		"email": "ryanpuj@ogmail.com",
		"password": "fdf"
	}
	`)
	testTable := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, message string, isError bool)
	}{
		"succes api call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("Update", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusOK, statusCode)
				require.NotEmpty(t, message)
				require.Equal(t, "succesfully updated", message)
				require.False(t, isError)
			},
		},
		"fail api call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, message)
				require.Equal(t, "got an error", message)
				require.True(t, isError)
			},
		},
		"bad json": {
			json:    wrongValidation,
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, statusCode int, message string, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, message)
				require.True(t, isError)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, _ := http.NewRequest(http.MethodPatch, "/user", bytes.NewReader(v.json))
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonResponse
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res.Message, res.Error)
		})
	}
}
