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
	"github.com/sodabitsdev/inventory-app/models"

	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"

	"database/sql"
	//"time"
	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

// Global variable
//var DB *gorm.DB

// main function to start the server
func main() {

	log.Debug("Debug message")
	log.Info("Info message")

	db, err := sql.Open("sqlite3", "./inventory.db")
	checkErr(err)

	// insert
	// stmt, err := db.Prepare("insert into price_books (barcode, product_description, price) values (?,?,?) ")
	// checkErr(err)

	// res, err := stmt.Exec(123, "ice cream", 2.99)
	// checkErr(err)

	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)

	// Query all
	priceBook, err := models.FindAllPriceBookItems(db)

	if err != nil {
		log.Error("FindAllPriceBookItems returned err", err)
	}

	log.Info("priceBook", priceBook)

	// Query by barcode
	priceBook1, err := models.FindPriceBookItemByBarcode(db, 123)
	if err != nil {
		log.Error("FindPriceBookItemByBarcode returned error", err)
	}

	log.Info("Query by barcode: ", priceBook1)

	// Insert
	priceBookInsert := &models.PriceBook{Barcode: 456, ProductDescription: "test product 456", Price: 456}
	err = models.InsertPriceBookItem(db, priceBookInsert)
	if err != nil {
		log.Error("InsertPriceBookItem returned an error: ", err)
	}

	// Update
	priceBookUpdate := &models.PriceBook{Barcode: 456, ProductDescription: "test product 456 - updated", Price: 10}

	err = models.UpdatePriceBookItem(db, priceBookUpdate)
	if err != nil {
		log.Error("UpdatePriceBookItem returned an error: ", err)
	}

	// dbConfig := config.BuildDBConfig()

	// db, err := gorm.Open(sqlite.Open(dbConfig.DBName), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("Error connecting to the database...will exit")
	// 	panic(err)
	// }

	// //defer db.Close()    // this method no longer exists in GORM 2.0

	// fmt.Printf("Established connection to this DB...%s\n", dbConfig.DBName)
	//println((db))

	// Create tables
	// db.AutoMigrate(&models.PriceBook{}, &models.Inventory{})

	// // Create records
	// db.Create(&models.PriceBook{Barcode: 123, ProductDescription: "New Product", Price: 1.23})
	// fmt.Println("Created record in PriceBook table")

	// db.Create(&models.Inventory{InventoryDate: time.Now(), Barcode: 123, ProductDescription: "Some product description", Price: 1.23, Quantity: 10})
	// fmt.Println("Created record in Inventory table")

	// var pb models.PriceBook
	// db.First(&pb)
	// fmt.Println(pb)

	// var pbs []models.PriceBook
	// db.Find(&pbs)
	// fmt.Println(pbs)

	//fmt.Println(result.RowsAffected)

	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	//     panic("failed to connect database")
	// }
	// defer db.Close()

	// var users []User
	// db.Find(&users)
	// fmt.Println("{}", users)

	// json.NewEncoder(w).Encode(users)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
