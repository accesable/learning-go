package product

import (
	"context"
	"database/sql"
	"math"
	"strconv"

	mysqlc "github.com/trann/e_commerce/internal/model"
	"github.com/trann/e_commerce/internal/types"
)

type Store struct {
	db      *sql.DB
	queries *mysqlc.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		queries: mysqlc.New(db),
	}
}

// GetProducts() ([]Product, error)
// GetProductByID() (*Product, error)
// DeleteProduct(id int) error
// UpdateProduct(id int) error
// GetProductsByID(id []int) ()

func (s *Store) GetProducts(ctx context.Context) ([]types.Product, error) {
	dbProducts, err := s.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	var products []types.Product
	for _, product := range dbProducts {
		item, err := convertDBProductToTypeProduct(product)
		if err != nil {
			return nil, err
		}
		products = append(products, item)
	}
	return products, nil
}
func (s *Store) GetProductByID(ctx context.Context, id int) (*types.Product, error) {
	return nil, nil
}

func (s *Store) DeleteProduct(ctx context.Context, id int) error {
	return nil
}
func (s *Store) UpdateProduct(ctx context.Context, id int) error {
	return nil
}

// roundToTwoDecimals rounds a float64 to two decimal places and ensures .00 rounding.
func roundToTwoDecimals(value float64) float64 {
	return math.Floor(value*100+0.5) / 100
}
func convertDBProductToTypeProduct(dbProduct mysqlc.Product) (types.Product, error) {
	price, err := strconv.ParseFloat(dbProduct.Price, 32)
	if err != nil {
		return types.Product{}, err
	}
	return types.Product{
		ID:          dbProduct.ID,
		Name:        dbProduct.Name,
		Image:       dbProduct.Image,
		Description: dbProduct.Description,
		CreatedAt:   dbProduct.Createdat,
		Price:       roundToTwoDecimals(price),
	}, nil
}
