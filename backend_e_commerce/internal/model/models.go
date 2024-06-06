// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package mysqlc

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type OrdersStatus string

const (
	OrdersStatusPending   OrdersStatus = "pending"
	OrdersStatusCompleted OrdersStatus = "completed"
	OrdersStatusCancelled OrdersStatus = "cancelled"
)

func (e *OrdersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrdersStatus(s)
	case string:
		*e = OrdersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrdersStatus: %T", src)
	}
	return nil
}

type NullOrdersStatus struct {
	OrdersStatus OrdersStatus
	Valid        bool // Valid is true if OrdersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrdersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrdersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrdersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrdersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrdersStatus), nil
}

type Order struct {
	ID        uint32
	Userid    uint32
	Total     string
	Status    OrdersStatus
	Address   string
	Createdat time.Time
}

type OrderItem struct {
	ID        uint32
	Orderid   uint32
	Productid uint32
	Quantity  int32
	Price     string
}

type Product struct {
	ID          uint32
	Name        string
	Description string
	Image       string
	Price       string
	Quantity    uint32
	Createdat   time.Time
}

type User struct {
	ID        uint32
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Createdat time.Time
}