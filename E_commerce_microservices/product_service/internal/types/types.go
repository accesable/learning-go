package types

import (
	"context"
	"database/sql"
	"time"
)

type CategoryStore interface{
  GetCategories(context context.Context) ([]Category,error)
}


type Category struct {
  ID        int32 `json:"id"`
  Name      string `json:"name"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

type Item struct {
	ID               int64
	Name             string
	CategoryID       sql.NullInt32
	ShortDescription sql.NullString
	OriginalPrice    sql.NullFloat64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
