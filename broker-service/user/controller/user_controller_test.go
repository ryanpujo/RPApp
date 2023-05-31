package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/spriigan/broker/authentication"
	"github.com/spriigan/broker/authentication/authmock"
	"github.com/spriigan/broker/infrastructure/router"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/controller"
	"github.com/spriigan/broker/user/controller/mocks"
	"github.com/spriigan/broker/user/domain"
	"github.com/spriigan/broker/user/user-proto/userpb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockClient *mocks.MockClient
	mockAuth   *authmock.MockAuth
	mux        *gin.Engine
	reqPayload = domain.User{
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	badPayload = domain.User{
		FirstName: "ry",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}

	userTest = &userpb.User{
		Id:        1,
		Firstname: "ryan",
		Lastname:  "pujo",
		Username:  "ryanpujo",
	}
	verr    = "min=3"
	uriVerr = "gt=0"
)

func TestMain(m *testing.M) {
	mockAuth = new(authmock.MockAuth)
	mockClient = new(mocks.MockClient)
	userController := controller.NewUserController(mockClient)
	mux = router.UserRoute(userController, &authentication.Authentication{AuthClient: mockAuth})
	os.Exit(m.Run())
}

func unAuthorizeAssert(t *testing.T, statusCode int, json response.JsonRes) {
	require.Equal(t, http.StatusUnauthorized, statusCode)
	require.NotEmpty(t, json.Error)
	require.Equal(t, "unauthorized", json.Error)
}

func TestCreate(t *testing.T) {
	payload, _ := json.Marshal(reqPayload)
	badJson, _ := json.Marshal(badPayload)
	testTable := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, json response.JsonRes)
	}{
		"user created": {
			json: payload,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("CreateUser", mock.Anything, mock.Anything).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusCreated, statusCode)
				require.NotEmpty(t, json)
				require.NotEmpty(t, json.User)
				require.Equal(t, reqPayload.Username, json.User.Username)
			},
		},
		"validation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, json)
				require.NotEmpty(t, json.Errors)
				require.Equal(t, verr, json.Errors["FirstName"])
			},
		},
		"failed to create": {
			json: payload,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("CreateUser", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, json.Error)
				require.Equal(t, "an error", json.Error)
			},
		},
		"unauthorize": {
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewReader(v.json))
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer jdbjrjrjt")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestGetById(t *testing.T) {
	testTable := map[string]struct {
		uri     string
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, json response.JsonRes)
	}{
		"got a user": {
			uri: "/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("GetById", mock.Anything, mock.Anything).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, json.Error)
				require.NotEmpty(t, json.User)
				require.Equal(t, userTest.Username, json.User.Username)
			},
		},
		"failed to get it": {
			uri: "/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, json.Error)
				require.Equal(t, "an error", json.Error)
			},
		},
		"validation error": {
			uri: "/0",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, json.Errors)
				require.Equal(t, uriVerr, json.Errors["Id"])
			},
		},
		"unauthorized": {
			uri: "/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodGet, v.uri, nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer kjndfnfj")
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
		uri     string
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, json response.JsonRes)
	}{
		"user deleted": {
			uri: "/delete/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("DeleteById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, json)
			},
		},
		"failed to delete": {
			uri: "/delete/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("DeleteById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, json.Error)
				require.Equal(t, "an error", json.Error)
			},
		},
		"validation error": {
			uri: "/delete/0",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, json.Errors)
				require.Equal(t, uriVerr, json.Errors["Id"])
			},
		},
		"unauthorized": {
			uri: "/delete/1",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodDelete, v.uri, nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer ndjnjf")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}
func TestUpdateById(t *testing.T) {
	payload, _ := json.Marshal(reqPayload)
	badJson, _ := json.Marshal(badPayload)
	testTable := map[string]struct {
		json    []byte
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, json response.JsonRes)
	}{
		"user updated": {
			json: payload,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("UpdateById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.Empty(t, json)
			},
		},
		"failed to update": {
			json: payload,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("UpdateById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, json.Error)
				require.Equal(t, "an error", json.Error)
			},
		},
		"validation error": {
			json: badJson,
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, json.Errors)
				require.Equal(t, verr, json.Errors["FirstName"])
			},
		},
		"unauthorized": {
			json: payload,
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
			req.Header.Add("Authorization", "Bearer ndjnjf")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}

func TestGetMany(t *testing.T) {
	users := &userpb.Users{
		Users: []*userpb.User{userTest},
	}
	testTable := map[string]struct {
		uri     string
		arrange func(t *testing.T)
		assert  func(t *testing.T, statusCode int, json response.JsonRes)
	}{
		"got users": {
			uri: "/users?page=1&limit=10",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("GetMany", mock.Anything, mock.Anything).Return(users, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusOK, statusCode)
				require.NotEmpty(t, json.Users)
				require.Equal(t, len(users.Users), len(json.Users))
			},
		},
		"failed to get users": {
			uri: "/users?page=1&limit=10",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
				mockClient.On("GetMany", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.NotEmpty(t, json.Error)
				require.Equal(t, "an error", json.Error)
			},
		},
		"validation error": {
			uri: "/users?page=1&limit=0",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(auth.Token{}, nil).Once()
			},
			assert: func(t *testing.T, statusCode int, json response.JsonRes) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.NotEmpty(t, json.Errors)
				require.Equal(t, uriVerr, json.Errors["Limit"])
			},
		},
		"unauthorized": {
			uri: "/users?page=1&limit=10",
			arrange: func(t *testing.T) {
				mockAuth.On("VerifyIDToken", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: unAuthorizeAssert,
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			req, err := http.NewRequest(http.MethodGet, v.uri, nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", "Bearer ndjnjf")
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			var res response.JsonRes
			json.NewDecoder(rr.Body).Decode(&res)

			v.assert(t, rr.Code, res)
		})
	}
}
