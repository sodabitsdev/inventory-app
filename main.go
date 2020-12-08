package main

import (

	/*
		The packages below are my packages.  I need to give a full path under the src directory.  So config and models are under github.com/sodabitsdev/inventory-app/
	*/

	/*
		database/sql

		https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html

		SQL tutorial in GO
		https://www.alexedwards.net/blog/practical-persistence-sql
	*/

	// "github.com/sodabitsdev/inventory-app/config"

	//"time"

	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sodabitsdev/inventory-app/models"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// main function to start the server
func main() {

	// Database handle
	var db *sqlx.DB

	// configure logger
	configureLogger()

	// load environment variables
	loadEnvVariables()
	env := os.Getenv("ENVIRONMENT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	log.Infoln("Loaded environment variables for.....", env)

	//db, err := sqlx.Open("sqlite3", "./inventory.db")
	connectString := dbUser + ":" + dbPassword + "@(" + dbHost + ")/" + dbDatabase
	log.Debugln("connectString ", connectString)

	db, err := sqlx.Connect("mysql", connectString)

	if err != nil {
		panic(err)
	}

	log.Infoln("Established connection to database ... ", dbHost, dbDatabase)

	// TESTING ....

	// Query all
	var priceBook []models.PriceBook
	priceBook, err = models.FindAllPriceBookItems(db)
	fmt.Println("printing results from FindAllPriceBookItems...")
	if priceBook != nil {
		fmt.Println(priceBook)
	}

	// priceBook, err = models.FindPriceBookItemByBarcode(db, "123")

	// if priceBook != nil {
	// 	fmt.Println("printing results from FindPriceBookItemByBarcode...")
	// 	fmt.Println(priceBook)
	// }

	// Insert
	priceBookInsert := &models.PriceBook{
		Barcode:            "456",
		ProductDescription: sql.NullString{String: "test product 789", Valid: true},
		Price:              sql.NullFloat64{Float64: 456},
	}

	err = models.InsertPriceBookItem(db, priceBookInsert)
	if err != nil {
		fmt.Println("InsertPriceBookItem returned an error: ", err)
	}

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	// Insert into Inventory table
	invDate, errdate := time.Parse(layoutISO, "2020-12-07")
	if errdate != nil {
		log.Error("date error : ", errdate)
	} else {
		log.Infoln("invDate: ", invDate)
	}

	inventoryInsert := &models.Inventory{
		InventoryDate:      invDate,
		Barcode:            "123",
		ProductDescription: sql.NullString{String: "testing", Valid: true},
		Price:              sql.NullFloat64{Float64: 100.00},
		Quantity:           10,
	}

	insertID, err := models.InsertInventory(db, inventoryInsert)
	if err != nil {
		log.Errorln("Error inserting inventory record: ", err)
	}

	log.Debugln("InsertInventory insertId ...", insertID)

	updateInventory := &models.Inventory{
		InventoryDate:      invDate,
		Barcode:            "123",
		ProductDescription: sql.NullString{String: "testing - UPDATED", Valid: true},
		Price:              sql.NullFloat64{Float64: 101.00},
		Quantity:           11,
	}

	updatedRows, err := models.UpdateInventory(db, updateInventory)
	if err != nil {
		log.Errorln("Error updateding inventory record: ", err)
	}

	log.Debugln("Updated Inventory record.  Records affected: ", updatedRows)

	/// TESTING end here ...

	// if err != nil {
	// 	log.Error("FindAllPriceBookItems returned err", err)
	// }

	// log.Info("priceBook", priceBook)

	// // Query by barcode
	// priceBook1, err := models.FindPriceBookItemByBarcode(db, 123)
	// if err != nil {
	// 	log.Error("FindPriceBookItemByBarcode returned error", err)
	// }

	// log.Info("Query by barcode: ", priceBook1)

	// // Insert
	// priceBookInsert := &models.PriceBook{Barcode: 456, ProductDescription: "test product 456", Price: 456}
	// err = models.InsertPriceBookItem(db, priceBookInsert)
	// if err != nil {
	// 	log.Error("InsertPriceBookItem returned an error: ", err)
	// }

	// // Update
	// priceBookUpdate := &models.PriceBook{Barcode: 456, ProductDescription: "test product 456 - updated", Price: 10}

	// err = models.UpdatePriceBookItem(db, priceBookUpdate)
	// if err != nil {
	// 	log.Error("UpdatePriceBookItem returned an error: ", err)
	// }

	// Query Inventories
	// inventories, err := models.FindAllInventories(db)
	// if err != nil {
	// 	log.Error("GetAllInventories returned error: ", err)
	// }
	// log.Info("inventories: ", inventories)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading environment variables...", err)
	}

}

func configureLogger() {

	//example found here: https://stackoverflow.com/questions/48971780/change-format-of-log-output-logrus/48972299

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	// print calling method in the log
	//log.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
	log.SetLevel(log.DebugLevel)

}
