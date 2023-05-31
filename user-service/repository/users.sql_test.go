package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/spriigan/RPApp/repository"
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
var userRepo *repository.Queries

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

	userRepo = repository.New(testDb)

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("cant clean up resources: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSql, err := os.ReadFile("./testdata/schema.sql")
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

func TestCreateUser(t *testing.T) {
	args := repository.CreateUserParams{
		FirstName: "ryan",
		LastName:  "pujo",
		Username:  "ryanpujo",
	}

	user, err := userRepo.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, int64(1), user.ID)
	require.Equal(t, "ryan", user.FirstName)
}

func TestGetById(t *testing.T) {
	user, err := userRepo.GetById(context.Background(), int64(1))
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, int64(1), user.ID)
	require.Equal(t, "ryan", user.FirstName)

	userNotFound, err := userRepo.GetById(context.Background(), int64(2))
	require.Error(t, err)
	require.Empty(t, userNotFound)
}

func TestGetMany(t *testing.T) {
	users, err := userRepo.GetMany(context.Background(), 3)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Equal(t, 1, len(users))
}

func TestUpdateById(t *testing.T) {
	args := repository.UpdateByIDParams{
		FirstName: "ryan",
		LastName:  "connor",
		Username:  "ryanconnor",
		ID:        1,
	}

	user, err := userRepo.GetById(context.Background(), int64(1))
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, int64(1), user.ID)
	require.NotEqual(t, args.LastName, user.LastName)

	err = userRepo.UpdateByID(context.Background(), args)
	require.NoError(t, err)

	user, err = userRepo.GetById(context.Background(), int64(1))
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.LastName, user.LastName)
}

func TestDeleteByid(t *testing.T) {
	err := userRepo.DeleteByID(context.Background(), int64(1))
	require.NoError(t, err)

	user, err := userRepo.GetById(context.Background(), int64(1))
	require.Error(t, err)
	require.Empty(t, user)
}
