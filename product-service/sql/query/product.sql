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

-- name: GetProductById :one
SELECT p.id, p.store_id, p.name, p.description, p.price, p.image_url, p.stock, p.category_id, s.store_name, c.name AS category FROM products p JOIN stores s ON p.store_id = s.id JOIN category c ON p.category_id = c.id WHERE p.id = $1;

-- name: GetProducts :many
SELECT p.id, p.store_id, p.name, p.description, p.price, p.image_url, p.stock, p.category_id, s.store_name, c.name AS category FROM products p JOIN stores s ON p.store_id = s.id JOIN category c ON p.category_id = c.id;

-- name: UpdateProduct :exec
UPDATE products SET store_id = $1, name = $2, description = $3, price = $4, image_url = $5, stock = $6, category_id = $7 WHERE id = $8;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id=$1;
