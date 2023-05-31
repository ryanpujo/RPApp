// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: users.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  "first_name",
  "last_name",
  "username"
) VALUES ($1, $2, $3)
RETURNING id, first_name, last_name, username, created_at
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.FirstName, arg.LastName, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const deleteByID = `-- name: DeleteByID :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteByID, id)
	return err
}

const getById = `-- name: GetById :one
SELECT "id", "first_name", "last_name", "username", "created_at" FROM users WHERE id = $1
`

func (q *Queries) GetById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const getMany = `-- name: GetMany :many
SELECT "id", "first_name", "last_name", "username", "created_at" FROM users LIMIT $1 OFFSET $2
`

type GetManyParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMany(ctx context.Context, arg GetManyParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getMany, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateByID = `-- name: UpdateByID :exec
UPDATE users SET first_name = $1, last_name = $2, username = $3 WHERE id = $4
`

type UpdateByIDParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	ID        int64  `json:"id"`
}

func (q *Queries) UpdateByID(ctx context.Context, arg UpdateByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateByID,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.ID,
	)
	return err
}
