package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/ryanpujo/product-service/internal/repository"
	"github.com/stretchr/testify/require"
)

const (
	host     = "localhost"
	port     = "5435"
	user     = "postgres"
	password = "postgres"
	dbName   = "products_test"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=20"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDb *sql.DB
var productRepo *repository.Queries

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

	productRepo = repository.New(testDb)

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("cant clean up resources: %s", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSql, err := os.ReadFile("./testdata/products schema.sql")
	if err != nil {
		log.Println(err)
		return err
	}
	querySql, err := os.ReadFile("./testdata/query.sql")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = testDb.Exec(string(tableSql))
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = testDb.Exec(string(querySql))
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

func TestCreateProduct(t *testing.T) {
	testProduct := repository.CreateProductParams{
		Name:        "MacBook",
		Description: "good product",
		Price:       "2000",
		ImageUrl:    "jdskjfd.com",
		StoreID:     1,
		CategoryID:  1,
		Stock:       30,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	createdProduct, err := productRepo.CreateProduct(ctx, testProduct)
	require.NoError(t, err)

	require.NotEmpty(t, createdProduct)
	require.Equal(t, testProduct.Name, createdProduct.Name)
	require.Equal(t, int64(1), createdProduct.ID)
	require.Equal(t, testProduct.StoreID, createdProduct.StoreID)
}
