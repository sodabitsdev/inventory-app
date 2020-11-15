package models

import (
	"gorm.io/gorm"
)

// GetAllPriceBooks do not use
func GetAllPriceBooks(priceBook *[]PriceBook) (err error) {
	if priceBook != nil {
		println("priceBook is not nil")
		return
	}

	println("priceBook is nil")
	return nil

}

// GetAllPrices returns a slice of PriceBook by reference or error
func GetAllPrices(db *gorm.DB, priceBook *[]PriceBook) (err error) {
	if err = db.Find(priceBook).Error; err != nil {
		return err
	}
	return nil
}
