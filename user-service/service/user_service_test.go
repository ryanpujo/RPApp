package service_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/spriigan/RPApp/repository"
	"github.com/spriigan/RPApp/service"
	"github.com/spriigan/RPApp/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	mockQuery   *mocks.MockQuery
	userService service.UserService
)

func TestMain(m *testing.M) {
	mockQuery = new(mocks.MockQuery)
	userService = service.NewUserService(mockQuery)
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	createArgs := repository.CreateUserParams{
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}

	createdUser := repository.User{
		ID:        1,
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual repository.User, err error)
	}{
		"user created": {
			arrange: func(t *testing.T) {
				mockQuery.On("CreateUser", mock.Anything, createArgs).Return(createdUser, nil).Once()
			},
			assert: func(t *testing.T, actual repository.User, err error) {
				require.NotEmpty(t, actual)
				require.NoError(t, err)
				require.Equal(t, createdUser, actual)
			},
		},
		"failed to create": {
			arrange: func(t *testing.T) {
				mockQuery.On("CreateUser", mock.Anything, createArgs).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual repository.User, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			actual, err := userService.CreateUser(context.Background(), createArgs)

			v.assert(t, actual, err)
		})
	}
}

func TestDeleteByID(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"user deleted": {
			arrange: func(t *testing.T) {
				mockQuery.On("DeleteByID", mock.Anything, int64(1)).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to delete": {
			arrange: func(t *testing.T) {
				mockQuery.On("DeleteByID", mock.Anything, int64(1)).Return(errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := userService.DeleteByID(context.Background(), int64(1))

			v.assert(t, err)
		})
	}
}

func TestGetById(t *testing.T) {
	foundUser := repository.User{
		ID:        1,
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual repository.User, err error)
	}{
		"got product": {
			arrange: func(t *testing.T) {
				mockQuery.On("GetById", mock.Anything, int64(1)).Return(foundUser, nil).Once()
			},
			assert: func(t *testing.T, actual repository.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, foundUser, actual)
			},
		},
		"did not get the product": {
			arrange: func(t *testing.T) {
				mockQuery.On("GetById", mock.Anything, int64(1)).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual repository.User, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			user, err := userService.GetById(context.Background(), int64(1))

			v.assert(t, user, err)
		})
	}
}

func TestGetMany(t *testing.T) {
	foundUser := repository.User{
		ID:        1,
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	users := []repository.User{foundUser}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual []repository.User, err error)
	}{
		"got products": {
			arrange: func(t *testing.T) {
				mockQuery.On("GetMany", mock.Anything, mock.Anything).Return(users, nil).Once()
			},
			assert: func(t *testing.T, actual []repository.User, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, actual)
				require.Equal(t, len(users), len(actual))
			},
		},
		"did not got any products": {
			arrange: func(t *testing.T) {
				mockQuery.On("GetMany", mock.Anything, mock.Anything).Return(nil, errors.New("an error")).Once()
			},
			assert: func(t *testing.T, actual []repository.User, err error) {
				require.Error(t, err)
				require.Empty(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			actual, err := userService.GetMany(context.Background(), int32(3))

			v.assert(t, actual, err)
		})
	}
}

func TestUpdateById(t *testing.T) {
	updateArgs := repository.UpdateByIDParams{
		ID:        1,
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"user updated": {
			arrange: func(t *testing.T) {
				mockQuery.On("UpdateByID", mock.Anything, updateArgs).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"failed to update": {
			arrange: func(t *testing.T) {
				mockQuery.On("UpdateByID", mock.Anything, updateArgs).Return(errors.New("an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := userService.UpdateByID(context.Background(), updateArgs)

			v.assert(t, err)
		})
	}
}
