-- name: CreateProduct :one
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
RETURNING *;