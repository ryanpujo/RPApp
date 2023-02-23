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
	"github.com/spriigan/broker/registry"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/user-proto/grpc/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type mockClient struct {
	mock.Mock
}

func (mc mockClient) RegisterUser(ctx context.Context, in *models.UserPayload, opts ...grpc.CallOption) (*models.UserId, error) {
	args := mc.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserId), args.Error(1)
}

func (mc mockClient) FindUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*models.Users, error) {
	return nil, nil
}

func (mc mockClient) FindByUsername(ctx context.Context, in *models.Username, opts ...grpc.CallOption) (*models.UserBio, error) {
	return nil, nil
}

func (mc mockClient) DeleteByUsername(ctx context.Context, in *models.Username, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (mc mockClient) Update(ctx context.Context, in *models.UserPayload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

var ac *adapters.AppController
var client *mockClient
var mux *gin.Engine
var res response.JsonResponse

func TestMain(m *testing.M) {
	client = new(mockClient)
	register := registry.New(client)
	ac = register.NewAppController()
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
		assert  func(t *testing.T, statusCode int, data interface{}, isError bool)
	}{
		"success api call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("RegisterUser", mock.Anything, mock.Anything).Return(&models.UserId{Id: int64(1)}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusOK, statusCode)
				require.False(t, isError)
				require.NotNil(t, data)
				require.Equal(t, float64(1), data.(float64))
			},
		},
		"failed call": {
			json: jsonReq,
			arrange: func(t *testing.T) {
				client.On("RegisterUser", mock.Anything, mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, data interface{}, isError bool) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Nil(t, data)
				require.True(t, isError)
			},
		},
		"wrong validation": {
			json:    wrongValidation,
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

			req, _ := http.NewRequest(http.MethodPost, "/user", bytes.NewReader(v.json))
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			_ = json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res.Data, res.Error)
		})
	}
}
