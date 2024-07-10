package types

import (
	"context"
	"time"
)

type CategoryStore interface {
	GetCategories(ctx context.Context) ([]Category, error)
	CreateCategory(ctx context.Context, payload *CreateCategoryPayload) (int64, error)
	DeleteCategory(ctx context.Context, id int) error
	GetCategoryById(ctx context.Context, id int) (Category, error)
	UpdateCategoryById(ctx context.Context, id int, name string) (int64, error)
}
type ItemStore interface {
	GetItems(ctx context.Context) ([]Item, error)
	CreateItem(ctx context.Context, payload *CreateItemPayload) (int64, error)
	// DeleteItem(ctx context.Context, id int) error
	DeleteItem(ctx context.Context, id int64) (int64, error)
}
type CreateItemPayload struct {
	Name             string  `json:"name"             validate:"required,min=3,max=128"`
	CategoryID       int32   `json:"categoryId"       validate:"required,number"`
	ShortDescription string  `json:"shortDescription"`
	OriginalPrice    float64 `json:"originalPrice"    validate:"required,number"`
}
type Category struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Item struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	CategoryID       int32     `json:"categoryId"`
	ShortDescription string    `json:"shortDescription"`
	OriginalPrice    float64   `json:"originalPrice"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type CreateCategoryPayload struct {
	Name string `json:"name" validate:"alpha,required,min=3,max=128"`
}
type UpdateCategoryPayload struct {
	ID   int32  `json:"id"          validate:"required,number"`
	Name string `json:"updatedName" validate:"required,alpha,min=3,max=128"`
}
