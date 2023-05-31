package client_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/spriigan/broker/user/grpc/client/mocks"
	"github.com/spriigan/broker/user/user-proto/userpb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockUserServer *mocks.MockUserServer
	client         userpb.UserServiceClient
	lis            *bufconn.Listener
	payload        = &userpb.UserPayload{
		Firstname: "ryan",
		Lastname:  "pujo",
		Username:  "ryanpujo",
	}
	userTest = &userpb.User{
		Firstname: "ryan",
		Lastname:  "pujo",
		Username:  "ryanpujo",
	}
)

func bufDialer(ctx context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	lis = bufconn.Listen(1024 * 1024)
	defer lis.Close()
	s := grpc.NewServer()
	defer s.Stop()
	mockUserServer = new(mocks.MockUserServer)
	userpb.RegisterUserServiceServer(s, mockUserServer)
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
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.User, err error)
	}{
		"user created": {
			arrange: func(t *testing.T) {
				mockUserServer.On("CreateUser", mock.Anything, mock.Anything).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, payload.Firstname, actual.Firstname)
			},
		},
		"failed to create": {
			arrange: func(t *testing.T) {
				mockUserServer.On("CreateUser", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
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

func TestGetById(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.User, err error)
	}{
		"got the user": {
			arrange: func(t *testing.T) {
				mockUserServer.On("GetById", mock.Anything, mock.Anything).Return(userTest, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, userTest.Firstname, actual.Firstname)
			},
		},
		"did not get the user": {
			arrange: func(t *testing.T) {
				mockUserServer.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
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

	users := &userpb.Users{
		Users: []*userpb.User{userTest},
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *userpb.Users, err error)
	}{
		"got the users": {
			arrange: func(t *testing.T) {
				mockUserServer.On("GetMany", mock.Anything, mock.Anything).Return(users, nil).Once()
			},
			assert: func(t *testing.T, actual *userpb.Users, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, len(users.Users), len(actual.Users))
			},
		},
		"failed to get users": {
			arrange: func(t *testing.T) {
				mockUserServer.On("GetMany", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
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

			actual, err := client.GetMany(context.Background(), &userpb.GetMAnyArgs{Limit: 3, Page: 3})

			v.assert(t, actual, err)
		})
	}
}

func TestUpdateById(t *testing.T) {
	payload.Id = 1
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"user updated": {
			arrange: func(t *testing.T) {
				mockUserServer.On("UpdateById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to update": {
			arrange: func(t *testing.T) {
				mockUserServer.On("UpdateById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			_, err := client.UpdateById(context.Background(), payload)

			v.assert(t, err)
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
				mockUserServer.On("DeleteById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to delete": {
			arrange: func(t *testing.T) {
				mockUserServer.On("DeleteById", mock.Anything, mock.Anything).Return(&emptypb.Empty{}, errors.New("an error")).Once()
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
