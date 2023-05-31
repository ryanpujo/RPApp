package controller_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/spriigan/RPApp/controller"
	"github.com/spriigan/RPApp/controller/mocks"
	"github.com/spriigan/RPApp/repository"
	"github.com/spriigan/RPApp/user-proto/userpb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var (
	client      userpb.UserServiceClient
	lis         *bufconn.Listener
	mockService *mocks.MockUserService
)
var userTest = repository.User{
	ID:        1,
	FirstName: "ryan",
	LastName:  "pujo",
	Username:  "ryanpujo",
}
var payload = &userpb.UserPayload{
	Firstname: "ryan",
	Lastname:  "pujo",
	Username:  "ryanpujo",
}

func bufDialer(ctx context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	lis = bufconn.Listen(1024 * 1024)
	defer lis.Close()
	s := grpc.NewServer()
	defer s.Stop()
	mockService = new(mocks.MockUserService)
	userServer := controller.NewUserController(mockService)
	userpb.RegisterUserServiceServer(s, userServer)
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client = userpb.NewUserServiceClient(conn)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	args := repository.CreateUserParams{
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.User, err error)
	}{
		"product created": {
			arrange: func(t *testing.T) {
				mockService.On("CreateUser", mock.Anything, args).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, userTest.Username, actual.Username)
			},
		},
		"failed to create": {
			arrange: func(t *testing.T) {
				mockService.On("CreateUser", mock.Anything, args).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			actual, err := client.CreateUser(context.Background(), payload)

			v.assert(t, actual, err)
		})
	}
}

func TestDeleteById(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"user deleted": {
			arrange: func(t *testing.T) {
				mockService.On("DeleteByID", mock.Anything, int64(1)).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to delete": {
			arrange: func(t *testing.T) {
				mockService.On("DeleteByID", mock.Anything, int64(1)).Return(errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.DeleteById(context.Background(), &userpb.UserId{Id: 1})

			v.assert(t, err)
		})
	}
}

func TestGetById(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.User, err error)
	}{
		"got user": {
			arrange: func(t *testing.T) {
				mockService.On("GetById", mock.Anything, int64(1)).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, userTest.Username, actual.Username)
			},
		},
		"dont get the user": {
			arrange: func(t *testing.T) {
				mockService.On("GetById", mock.Anything, int64(1)).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			actual, err := client.GetById(context.Background(), &userpb.UserId{Id: 1})

			v.assert(t, actual, err)
		})
	}
}

func TestGetMany(t *testing.T) {
	usersTest := []repository.User{userTest}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.Users, err error)
	}{
		"got users": {
			arrange: func(t *testing.T) {
				mockService.On("GetMany", mock.Anything, int32(3)).Return(usersTest, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.Users, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, len(usersTest), len(actual.Users))
			},
		},
		"did not get users": {
			arrange: func(t *testing.T) {
				mockService.On("GetMany", mock.Anything, int32(3)).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual *userpb.Users, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			users, err := client.GetMany(context.Background(), &userpb.GetMAnyArgs{Limit: 3, Page: 3})

			v.assert(t, users, err)
		})
	}
}

func TestUpdateById(t *testing.T) {
	arg := repository.UpdateByIDParams{
		ID:        1,
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	payload := userpb.UserPayload{
		Id:        1,
		Firstname: "ryan",
		Lastname:  "pujo",
		Username:  "ryanpujo",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"user updated": {
			arrange: func(t *testing.T) {
				mockService.On("UpdateByID", mock.Anything, arg).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to update": {
			arrange: func(t *testing.T) {
				mockService.On("UpdateByID", mock.Anything, arg).Return(errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.UpdateById(context.Background(), &payload)

			v.assert(t, err)
		})
	}
}
