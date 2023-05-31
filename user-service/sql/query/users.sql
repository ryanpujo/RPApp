-- name: CreateUser :one
INSERT INTO users (
  "first_name",
  "last_name",
  "username"
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetById :one
SELECT "id", "first_name", "last_name", "username", "created_at" FROM users WHERE id = $1;

-- name: GetMany :many
SELECT "id", "first_name", "last_name", "username", "created_at" FROM users LIMIT $1 OFFSET $2;

-- name: DeleteByID :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateByID :exec
UPDATE users SET first_name = $1, last_name = $2, username = $3 WHERE id = $4;