package models

import (
	"database/sql"
	"time"
)

// PriceBook is a structure for the table price_books
type PriceBook struct {
	Barcode            string          `db:"barcode" json:"barcode"`
	ProductDescription sql.NullString  `db:"product_description" json:"productDescription"`
	Price              sql.NullFloat64 `db:"price" json:"price"`
}

// Inventory is a structure for the table inventories
type Inventory struct {
	InventoryDate      time.Time       `db:"inventory_date" json:"inventoryDate"`
	Barcode            string          `db:"barcode" json:"barcode"`
	ProductDescription sql.NullString  `db:"product_description" json:"productDescription"`
	Price              sql.NullFloat64 `db:"price" json:"price"`
	Quantity           sql.NullInt64   `db:"quantity" json:"quantity"`
}
