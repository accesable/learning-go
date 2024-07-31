package order

import (
	"context"
	"time"

	"gorm.io/gorm"

	"trann/ecom/order_service/internal/models"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Preload("OrderDetails").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// CreateOrder creates a new order with order details in the database
func (s *Store) CreateOrder(ctx context.Context, order *models.Order) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for _, detail := range order.OrderDetails {
			detail.OrderID = order.ID
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Store) GetOrderById(ctx context.Context, id int) (models.Order, error) {
	var order models.Order
	if err := s.db.Preload("OrderDetails").First(&order, id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *Store) CreateOrderDetailsToOrder(
	ctx context.Context,
	orderId int,
	orderDetails []*models.OrderDetail,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingDetails []models.OrderDetail
		if err := tx.Where("order_id = ?", orderId).Find(&existingDetails).Error; err != nil {
			return err
		}

		// Create a map to track existing order details by ItemID
		existingDetailsMap := make(map[int64]*models.OrderDetail)
		for i := range existingDetails {
			detail := &existingDetails[i]
			existingDetailsMap[detail.ItemID] = detail
		}

		for _, detail := range orderDetails {
			if existingDetail, exists := existingDetailsMap[detail.ItemID]; exists {
				// If the detail already exists, increment the quantity
				existingDetail.Quantity += detail.Quantity
				if err := tx.Save(existingDetail).Error; err != nil {
					return err
				}
			} else {
				// If the detail does not exist, create a new one
				detail.OrderID = int64(orderId)
				if err := tx.Create(detail).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *Store) UpdateOrderTime(ctx context.Context, order *models.Order) error {
	order.LastUpdatedAt = time.Now()
	if err := s.db.WithContext(ctx).Save(&order).Error; err != nil {
		return err
	}
	return nil
}
