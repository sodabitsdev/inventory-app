package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

///////////////////////////////////////////////////////////////////////
// Table price_books functions
///////////////////////////////////////////////////////////////////////

// FindAllPriceBookItems returns all records from PriceBook table
// priceBook parameter is pass by reference so the return value are stored in it
func FindAllPriceBookItems(db *sqlx.DB) ([]PriceBook, error) {
	priceBook := []PriceBook{}
	err := db.Select(&priceBook, "select barcode, product_description, price from price_books")

	if err != nil {
		log.Error("Error when selecting all rows from price_books ", err)
		return nil, err
	}

	return priceBook, nil

	/*
		rows, err := db.Query("select barcode, product_description, price from price_books")
		var priceBook []PriceBook

		if err != nil {
			log.Error("error is not nil ", err)
			return nil, err
		}

		for rows.Next() {
			//pb := new(PriceBook)			// this didn't work
			pb := PriceBook{} // but this worked
			err := rows.Scan(&pb.Barcode, &pb.ProductDescription, &pb.Price)
			if err != nil {
				log.Error("scanning row got an error: ", err)
			}

			priceBook = append(priceBook, pb)

		}

		return priceBook, nil
	*/

	/*
		if priceBook != nil {
			// an element of slice cannot be accessed like priceBook[0]
			// so dereference it like below.
			// found solution here: https://stackoverflow.com/questions/38468258/why-is-indexing-on-the-slice-pointer-not-allowed-in-golang
			pb := (*priceBook)[0]
			db.First(pb)
		}
	*/

}

// FindPriceBookItemByBarcode return a slice of records, ideally one record
func FindPriceBookItemByBarcode(db *sqlx.DB, barcode string) ([]PriceBook, error) {

	var sql = "select barcode, product_description, price from price_books where barcode = ?"

	pb := []PriceBook{}
	err := db.Select(&pb, sql, barcode)
	if err != nil {
		return nil, err
	}

	return pb, nil

	/*
		var sql = "select barcode, product_description, price from price_books where barcode = $1"
		row := db.QueryRow(sql, barcode)

		var priceBook []PriceBook
		pb := PriceBook{}
		err := row.Scan(&pb.Barcode, &pb.ProductDescription, &pb.Price)
		if err != nil {
			log.Error("scanning row got an error: ", err)
		}

		priceBook = append(priceBook, pb)

		return priceBook, nil

	*/
}

// InsertPriceBookItem inserts one record in PriceBook table.  If successful
// return nil.  Not going to return rowsAffected since it is not supported by
// all database drivers
func InsertPriceBookItem(db *sqlx.DB, priceBook *PriceBook) error {
	log.Debugln("InsertPriceBookItem... from methods file")
	if len(priceBook.Barcode) == 0 {
		log.Error("barcode is not set to a value ... returning error")
		err := errors.New("barcode is not set to a real value")
		return err
	}

	var sql = "insert into price_books (barcode, product_description, price) " +
		"values(?, ?, ?)"

	db.MustExec(sql, priceBook.Barcode, priceBook.ProductDescription, priceBook.Price)

	return nil

	/*
		_, err := db.Exec(sql, priceBook.Barcode, priceBook.ProductDescription, priceBook.Price)

		if err != nil {
			log.Error("insert to price_books failed ", err)
			return err
		}

		log.Info("insert to price_books passed")
		return nil
	*/

}

// UpdatePriceBookItem updates one or more items in PriceBook table
// the primary key barcode must exist
func UpdatePriceBookItem(db *sql.DB, priceBook *PriceBook) error {

	var sql = `update price_books set product_description = $1, price = $2 
		 where barcode = $3`

	_, err := db.Exec(sql, priceBook.ProductDescription, priceBook.Price,
		priceBook.Barcode)

	if err != nil {
		log.Error("Update to price_books failed ", err)
		return nil
	}

	log.Debug("update to price_books passed")
	return nil

}

///////////////////////////////////////////////////////////////////////
// Table inventories functions
///////////////////////////////////////////////////////////////////////

// FindAllInventories returns all records from table inventories
func FindAllInventories(db *sql.DB) ([]Inventory, error) {
	var sql = `select inventory_date, barcode, product_description, price, quantity from inventories order by inventory_date`
	logrus.Debug("SQL: ", sql)

	rows, err := db.Query(sql)
	if err != nil {
		log.Error("Error from FindAllInventories: ", err)
		return nil, err
	}

	var inventory []Inventory

	for rows.Next() {
		inv := Inventory{}
		err := rows.Scan(&inv.InventoryDate, &inv.Barcode, &inv.ProductDescription, &inv.Price, &inv.Quantity)
		if err != nil {
			log.Error("scanning frow got an error: ", err)
		}
		inventory = append(inventory, inv)
	}

	return inventory, nil
}

// FindAllInventoriesByDate returns all records from table inventories that match on the InventoryDate column
func FindAllInventoriesByDate(db *sql.DB, inventoryDate *time.Time) ([]Inventory, error) {
	// var sql = `Select * from inventories where inventory_date = $1`
	// log.Debug("SQL: ", sql)

	// rows := db.QueryRow(sql, inventoryDate)

	// var inventory []Inventory
	// inv := Inventory{}
	// err := row.Scan()

	return nil, nil
}

// GetInventoryByBarcode returns all records from table inventories that match on the column barcode
func GetInventoryByBarcode(db *sql.DB, barcode string) ([]Inventory, error) {
	var sql = `Select * from inventories where barcode = $1`
	log.Debug("SQL: ", sql)

	return nil, nil
}

// InsertInventory inserts a record in table inventories
func InsertInventory(db *sql.DB, inventory *Inventory) error {
	log.Debug("InsertInventory method called")

	var sql = `Insert into inventories (inventory_date, 
										barcode, 
										product_description, 
										price, 
										quantity)
										values ($1, $2, $3, $4, $5)`
	log.Debug("SQL: ", sql)

	return nil
}

// UpdateInventory a record in table inventory where columns match on inventory_date and barcode
func UpdateInventory(db *sql.DB, inventory *Inventory) error {

	var sql = `Update inventories set product_description = $1, price = $2, 
				quantity = $3`
	log.Debug("SQL: ", sql)
	return nil
}
