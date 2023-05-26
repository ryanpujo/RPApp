package domain

import "database/sql"

type Product struct {
	ID          int64        `json:"id"`
	StoreID     int64        `json:"store_id" binding:"required"`
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Price       string       `json:"price" binding:"required,gt=0"`
	ImageUrl    string       `json:"image_url" binding:"required"`
	Stock       int32        `json:"stock" binding:"required"`
	CategoryID  int64        `json:"category_id" binding:"required"`
	CreatedAt   sql.NullTime `json:"created_at"`
}
