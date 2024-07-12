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
	GetItems(ctx context.Context, opts ...GetItemsOption) ([]Item, error)
	CreateItem(ctx context.Context, payload *CreateItemPayload) (int64, error)
	DeleteItem(ctx context.Context, id int64) (int64, error)
	GetItemImagesById(ctx context.Context, id int) ([]ItemImage, error)
	UploadImageToItemId(ctx context.Context, itemImage ItemImage) (int64, error)
	UpdateItemById(ctx context.Context, id int, updatePayload PartialUpdateItem) (int64, error)
}
type CreateItemPayload struct {
	Name             string  `json:"name"             validate:"required,min=3,max=128"`
	CategoryID       int32   `json:"categoryId"       validate:"required,number"`
	ShortDescription string  `json:"shortDescription"`
	OriginalPrice    float64 `json:"originalPrice"    validate:"required,number"`
}
type PartialUpdateItem struct {
	Name             *string   `json:"name,omitempty"`
	CategoryID       *int      `json:"categoryId,omitempty"`
	ShortDescription *string   `json:"shortDescription,omitempty"`
	OriginalPrice    *float64  `json:"originalPrice,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
type Category struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Item struct {
	ID               int64       `json:"id"`
	Name             string      `json:"name"`
	CategoryID       int32       `json:"categoryId"`
	CategoryName     string      `json:"categoryName,omitempty"`
	ShortDescription string      `json:"shortDescription"`
	OriginalPrice    float64     `json:"originalPrice"`
	ImgURLs          []ItemImage `json:"imgUrls,omitempty"`
	CreatedAt        time.Time   `json:"createdAt"`
	UpdatedAt        time.Time   `json:"updatedAt"`
}

type CreateCategoryPayload struct {
	Name string `json:"name" validate:"alpha,required,min=3,max=128"`
}
type UpdateCategoryPayload struct {
	ID   int32  `json:"id"          validate:"required,number"`
	Name string `json:"updatedName" validate:"required,alpha,min=3,max=128"`
}
type ItemImage struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"displayName"`
	ImageUrl    string `json:"imageUrl"`
	ItemID      int64  `json:"itemId"`
}
