package utils

import (
	"trann/ecom/order_service/internal/models"
	"trann/ecom/order_service/internal/types"
)

// ToOrderPayload converts an Order model to an OrderPayload
func ToOrderPayload(order models.Order) types.OrderPayload {
	orderDetails := make([]types.OrderDetailPayload, len(order.OrderDetails))
	for i, detail := range order.OrderDetails {
		orderDetails[i] = ToOrderDetailPayload(detail)
	}

	return types.OrderPayload{
		ID:            order.ID,
		CreatedAt:     order.CreatedAt,
		LastUpdatedAt: order.LastUpdatedAt,
		OrderDetails:  orderDetails,
		NumberOfItem:  len(orderDetails),
	}
}

// ToOrderDetailPayload converts an OrderDetail model to an OrderDetailPayload
func ToOrderDetailPayload(detail models.OrderDetail) types.OrderDetailPayload {
	return types.OrderDetailPayload{
		ID:            detail.ID,
		OrderID:       detail.OrderID,
		ItemID:        detail.ItemID,
		CreatedAt:     detail.CreatedAt,
		LastUpdatedAt: detail.LastUpdatedAt,
		Quantity:      detail.Quantity,
	}
}
