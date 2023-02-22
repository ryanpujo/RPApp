package repository_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/spriigan/RPApp/domain"
	repos "github.com/spriigan/RPApp/interface/repository"
	"github.com/spriigan/RPApp/usecases/repository"
	"github.com/stretchr/testify/require"
)

const (
	host     = "localhost"
	port     = "5435"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=20"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDb *sql.DB
var userRepo repository.UserRepository

func TestMain(m *testing.M) {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("cant connect to docker, make sure that docker is running: %s", err)
	}

	pool = p

	opt := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2-alpine",
		Env: []string{
			"POSTGRES_USER" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err = pool.RunWithOptions(&opt)
	if err != nil {
		log.Fatalf("couldn't start resources: %s", err)
	}

	if err = pool.Retry(func() error {
		var errDb error
		testDb, errDb = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if errDb != nil {
			log.Println(errDb)
		}
		return testDb.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("couldn't connect to database: %s", err)
	}

	err = createTables()
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("cant create table: %s", err)
	}

	userRepo = repos.NewUserRepository(testDb)

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("cant clean up resources: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSql, err := os.ReadFile("./testdata/user.sql")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = testDb.Exec(string(tableSql))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func TestPingDb(t *testing.T) {
	err := testDb.Ping()
	require.NoError(t, err)
}

func TestCreate(t *testing.T) {
	payload := domain.UserPayload{
		Fname:    "ryan",
		Lname:    "pujo",
		Username: "ryanpujo",
		Email:    "ryanpujo@gmail.com",
		Password: "oke",
	}

	id, err := userRepo.Create(&payload)
	require.NoError(t, err)
	require.Equal(t, 1, id)
}

func TestFindUsers(t *testing.T) {
	payload := domain.UserPayload{
		Fname:    "ryan",
		Lname:    "pujo",
		Username: "ryanpujo1",
		Email:    "ryanpujo1@gmail.com",
		Password: "oke",
	}
	_, err := userRepo.Create(&payload)
	require.NoError(t, err)

	actual, err := userRepo.FindUsers()
	require.NoError(t, err)
	require.NotEmpty(t, actual)
	require.Equal(t, 2, len(actual))
}

func TestFindByUsername(t *testing.T) {
	user, err := userRepo.FindByUsername("ryanpujo1")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Email, "ryanpujo1@gmail.com")
	user, err = userRepo.FindByUsername("oke")
	require.Error(t, err)
	require.EqualError(t, err, repos.ErrNoUserFound.Error())
	require.Nil(t, user)
}

func TestDeleteByUsername(t *testing.T) {
	err := userRepo.DeleteByUsername("ryanpujo1")
	require.NoError(t, err)
	user, err := userRepo.FindByUsername("ryanpujo1")
	require.Error(t, err)
	require.EqualError(t, err, repos.ErrNoUserFound.Error())
	require.Nil(t, user)
}

func TestUpdate(t *testing.T) {
	payload := domain.UserPayload{
		Id:       1,
		Fname:    "ryan",
		Lname:    "conor",
		Username: "ryanpujo",
		Email:    "ryanpujo@gmail.com",
		Password: "oke",
	}
	err := userRepo.Update(payload)
	require.NoError(t, err)
	user, err := userRepo.FindByUsername("ryanpujo")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, payload.Lname, user.Lname)
}
