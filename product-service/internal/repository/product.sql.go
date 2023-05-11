// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: product.sql

package repository

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  store_id,
  name, 
  description,
  price,
  image_url,
  stock,
  category_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, store_id, name, description, price, image_url, stock, category_id, created_at
`

type CreateProductParams struct {
	StoreID     int32  `json:"store_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImageUrl    string `json:"image_url"`
	Stock       int32  `json:"stock"`
	CategoryID  int32  `json:"category_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.StoreID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ImageUrl,
		arg.Stock,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageUrl,
		&i.Stock,
		&i.CategoryID,
		&i.CreatedAt,
	)
	return i, err
}