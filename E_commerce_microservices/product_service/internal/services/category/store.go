package category

import (
	"context"

	mysqlc "trann/ecom/product_services/internal/model"
	"trann/ecom/product_services/internal/types"
)

type Store struct {
	queries *mysqlc.Queries
}

func NewStore(queries *mysqlc.Queries) *Store {
	return &Store{
		queries: queries,
	}
}

func (s *Store) GetCategories(ctx context.Context) ([]types.Category, error) {
	db, err := s.queries.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	var categories []types.Category
	for _, v := range db {
		categories = append(categories, convertDBCategoryToPayloadCategory(&v))
	}
	return categories, nil
}

func (s *Store) CreateCategory(
	ctx context.Context,
	payload *types.CreateCategoryPayload,
) (int64, error) {
	result, err := s.queries.CreateCategory(ctx, payload.Name)
	if err != nil {
		return 0, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertedId, nil
}

// GetCategories(context context.Context) ([]Item,error)
func convertDBCategoryToPayloadCategory(v *mysqlc.Category) types.Category {
	return types.Category{
		ID:        v.ID,
		Name:      v.Name,
		CreatedAt: v.CreatedAt.Time,
		UpdatedAt: v.UpdatedAt.Time,
	}
}

func (s *Store) DeleteCategory(ctx context.Context, id int) error {
	if err := s.queries.DeleteCategory(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCategoryById(ctx context.Context, id int) (types.Category, error) {
	db, err := s.queries.GetCategoryByID(ctx, int32(id))
	if err != nil {
		return types.Category{}, err
	}
	return convertDBCategoryToPayloadCategory(&db), nil
}

func (s *Store) UpdateCategoryById(ctx context.Context, id int, updatedName string) (int64, error) {
	result, err := s.queries.UpdateCategory(ctx, mysqlc.UpdateCategoryParams{
		ID:   int32(id),
		Name: updatedName,
	})
	if err != nil {
		return 0, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return changedRows, nil
}
