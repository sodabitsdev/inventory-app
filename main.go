package main

import (
	"fmt"

	"github.com/sodabitsdev/inventory-app/config"
	"github.com/sodabitsdev/inventory-app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global variable
//var DB *gorm.DB

// main function to start the server
func main() {

	dbConfig := config.BuildDBConfig()

	db, err := gorm.Open(sqlite.Open(dbConfig.DBName), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to the database...will exit")
		panic(err)
	}

	//defer db.Close()    // this method no longer exists in GORM 2.0

	fmt.Printf("Established connection to this DB...%s\n", dbConfig.DBName)
	//println((db))

	// Create tables
	// db.AutoMigrate(&models.PriceBook{}, &models.Inventory{})

	// // Create records
	// db.Create(&models.PriceBook{Barcode: 123, ProductDescription: "New Product", Price: 1.23})
	// fmt.Println("Created record in PriceBook table")

	// db.Create(&models.Inventory{InventoryDate: time.Now(), Barcode: 123, ProductDescription: "Some product description", Price: 1.23, Quantity: 10})
	// fmt.Println("Created record in Inventory table")

	var pb models.PriceBook
	db.First(&pb)
	fmt.Println(pb)

	var pbs []models.PriceBook
	db.Find(&pbs)
	fmt.Println(pbs)

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
