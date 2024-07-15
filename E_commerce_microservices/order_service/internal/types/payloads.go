package types

import "time"

// OrderPayload represents the JSON payload for an order
type OrderPayload struct {
	ID            int64                `json:"id"`
	CreatedAt     time.Time            `json:"createdAt"`
	LastUpdatedAt time.Time            `json:"lastUpdatedAt"`
	NumberOfItem  int                  `json:"numberOfItem"`
	OrderDetails  []OrderDetailPayload `json:"orderDetails"`
}

// OrderDetailPayload represents the JSON payload for an order detail
type OrderDetailPayload struct {
	ID            int64     `json:"id"`
	OrderID       int64     `json:"orderId"`
	ItemID        int64     `json:"itemId"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	Quantity      int       `json:"quantity"`
}

// CreateOrderPayload represents the JSON payload for creating a new order
type CreateOrderPayload struct {
	OrderDetails []CreateOrderDetailPayload `json:"orderDetails" binding:"required"`
}

type AddOrderDetailsRequest struct {
	OrderDetails []CreateOrderDetailPayload `json:"orderDetails" binding:"required"`
}

// CreateOrderDetailPayload represents the JSON payload for creating a new order detail
type CreateOrderDetailPayload struct {
	ItemID   int64 `json:"itemId"   binding:"required"`
	Quantity int   `json:"quantity" binding:"required"`
}
