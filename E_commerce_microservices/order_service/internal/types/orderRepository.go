package types

import (
	"context"

	"trann/ecom/order_service/internal/models"
)

type OrderRepository interface {
	GetOrders(context.Context) ([]models.Order, error)
	CreateOrder(ctx context.Context, order *models.Order) error
	GetOrderById(ctx context.Context, id int) (models.Order, error)
	CreateOrderDetailsToOrder(
		ctx context.Context,
		orderId int,
		orderDetails []*models.OrderDetail,
	) error
	UpdateOrderTime(ctx context.Context, order *models.Order) error
}
