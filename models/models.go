package models

import (
	"time"

	"gorm.io/gorm"
)

// PriceBook is a structure for the table price_books
type PriceBook struct {
	gorm.Model
	Barcode            uint    `gorm:"primaryKey" json:"barcode"`
	ProductDescription string  `json:"productDescription"`
	Price              float32 `gorm:"type:numeric(15,2)" json:"price"`
}

// Inventory is a structure for the table inventories
type Inventory struct {
	gorm.Model
	InventoryDate      time.Time `gorm:"type:date;primaryKey" json:"inventoryDate"`
	Barcode            uint      `gorm:"primaryKey" json:"barcode"` /* Barcode may or may not exist */
	ProductDescription string    `json:"productDescription"`
	Price              float32   `sql:"type:decimal(10,2)" json:"price"`
	Quantity           uint
}
