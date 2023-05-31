package registry

import (
	"database/sql"

	"github.com/spriigan/RPApp/user-proto/grpc/models"
)

type Registry interface {
	NewUserServer() models.UserServiceServer
}

type registry struct {
	DB *sql.DB
}

func New(db *sql.DB) *registry {
	return &registry{DB: db}
}
