package models

import (
	"time"
)

type Order struct {
	ID            int64         `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time     `gorm:"not null"`
	LastUpdatedAt time.Time     `gorm:"not null"`
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID"`
	Status        string        `gorm:"not null"`
}

// OrderDetail represents the "order_details" table in the database
type OrderDetail struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	OrderID       int64     `gorm:"not null"`
	ItemID        int64     `gorm:"not null"`
	CreatedAt     time.Time `gorm:"not null"`
	LastUpdatedAt time.Time `gorm:"not null"`
	Quantity      int       `gorm:"not null"`
	Order         Order     `gorm:"foreignKey:OrderID"`
	// Product       Product   `gorm:"foreignKey:ProductID"` // Assuming you have a Product model
}
