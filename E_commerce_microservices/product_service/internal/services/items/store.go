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

func (s *Store) GetItemImagesById(ctx context.Context, id int) ([]types.ItemImage, error) {
	return nil, nil
}

func (s *Store) UploadImageToItemId(ctx context.Context, itemImage types.ItemImage) (int64, error) {
	res, err := s.queries.InsertImage(ctx, mysqlc.InsertImageParams{
		DisplayName: sql.NullString{String: itemImage.DisplayName, Valid: true},
		ImageUrl:    itemImage.ImageUrl,
		ItemID:      itemImage.ItemID,
	})
	if err != nil {
		return 0, err
	}
	imgId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return imgId, nil
}

func (s *Store) UpdateItemById(
	ctx context.Context,
	id int,
	updatePayload types.PartialUpdateItem,
) (int64, error) {
	result, err := s.queries.PartialUpdateItem(ctx, mysqlc.PartialUpdateItemParams{
		ID: int64(id),
		Name: func() sql.NullString {
			if updatePayload.Name != nil {
				return sql.NullString{String: *updatePayload.Name, Valid: true}
			}
			return sql.NullString{Valid: false}
		}(),
		CategoryID: func() sql.NullInt32 {
			if updatePayload.CategoryID != nil {
				return sql.NullInt32{Int32: int32(*updatePayload.CategoryID), Valid: true}
			}
			return sql.NullInt32{Valid: false}
		}(),
		ShortDescription: func() sql.NullString {
			if updatePayload.ShortDescription != nil {
				return sql.NullString{String: *updatePayload.ShortDescription, Valid: true}
			}
			return sql.NullString{Valid: false}
		}(),
		OriginalPrice: func() sql.NullFloat64 {
			if updatePayload.OriginalPrice != nil {
				return sql.NullFloat64{Float64: *updatePayload.OriginalPrice, Valid: true}
			}
			return sql.NullFloat64{Valid: false}
		}(),
	})
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (s *Store) GetItems(ctx context.Context, opts ...types.GetItemsOption) ([]types.Item, error) {
	options := types.NewGetItemsOptions()
	for _, opt := range opts {
		opt(options)
	}
	itemsDb, err := s.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var items []types.Item
	for _, v := range itemsDb {
		item := convertToItemPayload((*mysqlc.Item)(&v))
		if options.ShowCategoryName {
			// get category name here
			category, err := s.queries.GetCategoryByID(ctx, item.CategoryID)
			if err != nil {
				return nil, err
			}
			item.CategoryName = category.Name
		}
		if options.IncludeImgURLs {
			// get image urls of the item
			imgUrls, err := s.queries.ListImagesByItemId(ctx, item.ID)
			if err != nil {
				return nil, err
			}
			for _, v1 := range imgUrls {
				url := convertToImgPayload((*mysqlc.ItemImage)(&v1))
				item.ImgURLs = append(item.ImgURLs, url)
			}
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Store) CreateItem(ctx context.Context, payload *types.CreateItemPayload) (int64, error) {
	res, err := s.queries.InsertProduct(ctx, mysqlc.InsertProductParams{
		Name:             sql.NullString{String: payload.Name, Valid: true},
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

func convertToItemPayload(item *mysqlc.Item) types.Item {
	return types.Item{
		ID:               item.ID,
		Name:             item.Name.String,
		CategoryID:       item.CategoryID.Int32,
		ShortDescription: item.ShortDescription.String,
		OriginalPrice:    item.OriginalPrice.Float64,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.CreatedAt,
	}
}

func convertToImgPayload(i *mysqlc.ItemImage) types.ItemImage {
	return types.ItemImage{
		ID:          i.ID,
		DisplayName: i.DisplayName.String,
		ImageUrl:    i.ImageUrl,
		ItemID:      i.ItemID,
	}
}
