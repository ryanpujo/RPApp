package repository

import (
	"context"
	"database/sql"
)

type QueriesInterface interface {
	GetProductById(ctx context.Context, id int64) (GetProductByIdRow, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	GetProducts(ctx context.Context) ([]GetProductsRow, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	WithTx(tx *sql.Tx) *Queries
}
