package domain

import (
	"database/sql"
	"time"
)

// Category represents a product category
type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// Product represents a product
type Product struct {
	ID            int32          `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	OriginalPrice float64        `json:"original_price"`
	CategoryID    int32          `json:"category_id"`
	CreatedAt     time.Time      `json:"created_at"`
	LastUpdatedAt time.Time      `json:"last_updated_at"`
}

// ProductImage represents an image associated with a product
type ProductImage struct {
	ID        int32  `json:"id"`
	ImageUrl  string `json:"image_url"`
	ProductID int32  `json:"product_id"`
}
