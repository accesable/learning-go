package items

import (
	"context"
	"database/sql"

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

func (s *Store) GetItems(ctx context.Context) ([]types.Item, error) {
	itemsDb, err := s.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var items []types.Item
	for _, v := range itemsDb {
		items = append(items, convertToPayload((*mysqlc.Item)(&v)))
	}
	return items, nil
}

func (s *Store) CreateItem(ctx context.Context, payload *types.CreateItemPayload) (int64, error) {
	res, err := s.queries.InsertProduct(ctx, mysqlc.InsertProductParams{
		Name:             payload.Name,
		CategoryID:       sql.NullInt32{Int32: payload.CategoryID, Valid: true},
		OriginalPrice:    sql.NullFloat64{Float64: payload.OriginalPrice, Valid: true},
		ShortDescription: sql.NullString{String: payload.ShortDescription, Valid: true},
	})
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) DeleteItem(ctx context.Context, id int64) (int64, error) {
	res, err := s.queries.DeleteProduct(ctx, id)
	if err != nil {
		return 0, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedRows, nil
}

func convertToPayload(item *mysqlc.Item) types.Item {
	return types.Item{
		ID:               item.ID,
		Name:             item.Name,
		CategoryID:       item.CategoryID.Int32,
		ShortDescription: item.ShortDescription.String,
		OriginalPrice:    item.OriginalPrice.Float64,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.CreatedAt,
	}
}
