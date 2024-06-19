package items

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
