package domain

import "database/sql"

type User struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"first_name" binding:"required,min=3"`
	LastName  string       `json:"last_name" binding:"required,min=3"`
	Username  string       `json:"username" binding:"required,min=3"`
	CreatedAt sql.NullTime `json:"created_at"`
}
