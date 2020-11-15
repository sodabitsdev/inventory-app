package config

import (
	"gorm.io/gorm"
)

var DB *gorm.DB // global variable

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// TODO:  Get these variables from an environment file
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "",
		Password: "",
		DBName:   "inventory.db",
	}

	return &dbConfig
}

// func DbURL(dbConfig *DBConfig) string {
// 	// github.com/mattn/go-sqlite3
// 	return db, err := gorm.Open(sqlite.Open(dbConfig.DBName), &gorm.Config{})

// }

// Sum up two numbers
func Sum(x int, y int) int {
	return x + y

}
