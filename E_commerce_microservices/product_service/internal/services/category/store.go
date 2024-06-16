package category

import (
	"context"
	"database/sql"
	mysqlc "trann/ecom/product_services/internal/model"
	"trann/ecom/product_services/internal/types"
)

type Store struct {
  db *sql.DB
  queries mysqlc.Queries
}

func NewStore(db *sql.DB) *Store {
  return &Store{
    db: db,
    queries: *mysqlc.New(db),
  } 
}

func (s *Store)GetCategories(ctx context.Context) ([]types.Category,error){
  db,err := s.queries.ListCategories(ctx)
  if err != nil {
    return nil,err
  }
  var categories []types.Category
  for _, v := range db {
    categories = append(categories, convertDBCategoryToPayloadCategory(&v))
  }
  db = nil;
  return categories,nil
}
//  GetCategories(context context.Context) ([]Item,error)
func convertDBCategoryToPayloadCategory(v *mysqlc.Category) (types.Category){
  return types.Category{
    ID: v.ID,
    Name: v.Name,
    CreatedAt: v.CreatedAt.Time,
    UpdatedAt: v.UpdatedAt.Time,
  }
}
