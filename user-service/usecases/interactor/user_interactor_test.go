package interactor_test

import (
	"errors"
	"os"
	"testing"

	"github.com/spriigan/RPApp/usecases/interactor"
	"github.com/spriigan/RPApp/user-proto/grpc/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	mock.Mock
}

func (in *mockUserRepo) Create(user *models.UserPayload) (int, error) {
	args := in.Called(user)
	return args.Int(0), args.Error(1)
}

func (in *mockUserRepo) FindUsers() (*models.Users, error) {
	args := in.Called()
	arg1 := args.Get(0)
	if arg1 == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Users), args.Error(1)
}

func (in *mockUserRepo) FindByUsername(username string) (*models.User, error) {
	args := in.Called()
	arg1 := args.Get(0)
	if arg1 == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (in *mockUserRepo) DeleteByUsername(username string) error {
	args := in.Called()
	return args.Error(0)
}

func (in *mockUserRepo) Update(user *models.UserPayload) error {
	args := in.Called()
	return args.Error(0)
}

var userInteractor interactor.UserInteractor
var mockRepo *mockUserRepo

func TestMain(m *testing.M) {
	mockRepo = new(mockUserRepo)
	userInteractor = interactor.NewUserInteractor(mockRepo)
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, id int, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				mockRepo.On("Create", mock.Anything).Return(1, nil).Once()
			},
			assert: func(t *testing.T, id int, err error) {
				require.NoError(t, err)
				require.Equal(t, 1, id)
			},
		},
		"fail call": {
			arrange: func(t *testing.T) {
				mockRepo.On("Create", mock.Anything).Return(0, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, id int, err error) {
				require.Error(t, err)
				require.Zero(t, id)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			result, err := userInteractor.Create(&models.UserPayload{})

			v.assert(t, result, err)
		})
	}
}

func TestFindUsers(t *testing.T) {
	users := &models.Users{
		User: []*models.UserBio{
			{},
			{},
		},
	}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *models.Users, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				mockRepo.On("FindUsers").Return(users, nil).Once()
			},
			assert: func(t *testing.T, actual *models.Users, err error) {
				require.NoError(t, err)
				require.Equal(t, users, actual)
				require.Equal(t, len(users.User), len(actual.User))
			},
		},
		"fail call": {
			arrange: func(t *testing.T) {
				mockRepo.On("FindUsers", mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, actual *models.Users, err error) {
				require.Error(t, err)
				require.Zero(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			result, err := userInteractor.FindUsers()

			v.assert(t, result, err)
		})
	}
}

func TestFindByUsername(t *testing.T) {
	user := &models.User{Fname: "dabi"}
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, actual *models.User, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				mockRepo.On("FindByUsername", mock.Anything).Return(user, nil).Once()
			},
			assert: func(t *testing.T, actual *models.User, err error) {
				require.NoError(t, err)
				require.NotNil(t, actual)
				require.Equal(t, user.Fname, actual.Fname)
			},
		},
		"fail call": {
			arrange: func(t *testing.T) {
				mockRepo.On("FindByUsername", mock.Anything).Return(nil, errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, actual *models.User, err error) {
				require.Error(t, err)
				require.Nil(t, actual)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			result, err := userInteractor.FindByUsername("")

			v.assert(t, result, err)
		})
	}
}

func TestDeleteByUsername(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				mockRepo.On("DeleteByUsername", mock.Anything).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"fail call": {
			arrange: func(t *testing.T) {
				mockRepo.On("DeleteByUsername", mock.Anything).Return(errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := userInteractor.DeleteByUsername("")

			v.assert(t, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	testTable := map[string]struct {
		arrange func(t *testing.T)
		assert  func(t *testing.T, err error)
	}{
		"succes call": {
			arrange: func(t *testing.T) {
				mockRepo.On("Update", mock.Anything).Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		"fail call": {
			arrange: func(t *testing.T) {
				mockRepo.On("Update", mock.Anything).Return(errors.New("got an error")).Once()
			},
			assert: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for k, v := range testTable {
		t.Run(k, func(t *testing.T) {
			v.arrange(t)

			err := userInteractor.Update(&models.UserPayload{})

			v.assert(t, err)
		})
	}
}
