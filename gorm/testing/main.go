package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/ecom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Query to get all items
	// var result []map[string]interface{}
	// if err := db.Table("items").Find(&result).Error; err != nil {
	// 	panic(err)
	// }
	var reqs []AddOrderDetailRequest
	reqs = append(reqs, AddOrderDetailRequest{
		ItemId:          1,
		SelectedOptions: []int{1, 2, 3}, Quantity: 2,
	})
	var order Order
	// var detailMap map[orderDetailTup]int
	db.Preload("OrderDetails").Find(&order, "2")
	for _, v := range order.OrderDetails {
		fmt.Printf("%v - %v \n", v.ItemID, v.Quantity)
	}
}

type AddOrderDetailRequest struct {
	ItemId          int
	SelectedOptions []int
	Quantity        int
}
type orderDetailTup struct {
	ItemId   int
	optionId int
}
type Order struct {
	ID            int64         `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time     `gorm:"not null"`
	LastUpdatedAt time.Time     `gorm:"not null"`
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID"`
	Status        string        `gorm:"not null"`
}

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
type OptionGroup struct {
	ID                 int       `gorm:"primaryKey;autoIncrement"`
	Name               string    `gorm:"type:nvarchar(255);not null"`
	MaxSelectedOptions int       `gorm:"default:1;not null"`
	MinSelectedOptions int       `gorm:"default:1;not null"`
	CreatedAt          time.Time `gorm:"type:datetime(6);not null"`
	LastUpdatedAt      time.Time `gorm:"type:datetime(6);not null"`
	Options            []Option  `gorm:"foreignKey:GroupID"` // One-to-Many relationship
}

// Option represents the options table
type Option struct {
	ID            int          `gorm:"primaryKey;autoIncrement"`
	Name          string       `gorm:"type:nvarchar(128);not null"`
	Description   string       `gorm:"type:longtext"`
	GroupID       int          `gorm:"not null"`
	CreatedAt     time.Time    `gorm:"type:datetime(6);not null"`
	LastUpdatedAt time.Time    `gorm:"type:datetime(6);not null"`
	Group         OptionGroup  `gorm:"foreignKey:GroupID"`  // Belongs to OptionGroup
	OptionItems   []OptionItem `gorm:"foreignKey:OptionID"` // One-to-Many relationship
}

// OptionItem represents the options_items table
type OptionItem struct {
	ID                 int       `gorm:"primaryKey;autoIncrement"`
	PriceModifierValue float64   `gorm:"type:float"`
	ItemID             int64     `gorm:"not null"`
	OptionID           int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"type:datetime(6);not null"`
	LastUpdatedAt      time.Time `gorm:"type:datetime(6);not null"`
	Option             Option    `gorm:"foreignKey:OptionID"` // Belongs to Option
}

// SelectedOptionDetail represents the selected_options_details table
type SelectedOptionDetail struct {
	ID           int         `gorm:"primaryKey;autoIncrement"`
	DetailID     int         `gorm:"not null"`
	OptionItemID int         `gorm:"not null"`
	Detail       OrderDetail `gorm:"foreignKey:DetailID"`     // Belongs to OrderDetail (assuming OrderDetail struct exists)
	OptionItem   OptionItem  `gorm:"foreignKey:OptionItemID"` // Belongs to OptionItem
}
